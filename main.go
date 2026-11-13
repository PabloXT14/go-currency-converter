package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/pabloxt14/go-currency-converter/models"
)

const (
	ratesFilePath = "data/rates.json"
	minArgsLength = 2
	usageMessage  = "uso: ./converter [valor_em_reais] [moeda_destino]. Ex: ./converter 100 USD"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	amount, currency, err := parseArgs(args)
	if err != nil {
		return err
	}

	rates, err := loadRates(ratesFilePath)
	if err != nil {
		return err
	}

	rate, ok := rates.Rates[currency]
	if !ok {
		return fmt.Errorf(
			"moeda inválida %q. Moedas disponíveis: %s",
			currency,
			strings.Join(availableCurrencies(rates.Rates), ", "),
		)
	}

	fmt.Printf("%.2f BRL = %.2f %s\n", amount, amount*rate, currency)

	return nil
}

func parseArgs(args []string) (amount float64, currency string, err error) {
	if len(args) < minArgsLength {
		return 0, "", fmt.Errorf(usageMessage)
	}

	amount, err = strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, "", fmt.Errorf("valor inválido %q: informe um número válido", args[0])
	}

	return amount, strings.ToUpper(args[1]), nil
}

// loadRates usa json.NewDecoder para decodificar direto do arquivo, sem carregar todo o conteúdo em memória (como os.ReadFile faria).
func loadRates(path string) (models.Rates, error) {
	var rates models.Rates

	file, err := os.Open(path)
	if err != nil {
		return rates, fmt.Errorf("abrindo arquivo de taxas: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&rates); err != nil {
		return rates, fmt.Errorf("decodificando taxas: %w", err)
	}

	return rates, nil
}

func availableCurrencies(rates map[string]float64) []string {
	return slices.Sorted(maps.Keys(rates))
}

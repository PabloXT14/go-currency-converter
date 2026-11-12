package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pabloxt14/go-currency-converter/models"
)

func main() {
	path := filepath.Join("data/rates.json")

	// Read the file data/rates.json slurping it's entire content into memory (not recommended for large files)
	// file, err := os.ReadFile(path)

	// check(err)

	// data := map[string]models.Rates{}

	// err = json.Unmarshal(file, &data)

	// check(err)

	// Json.NewDecoder() is more performant => Lê e decodifica diretamente de um io.Reader, ou seja, linha por linha, sendo útil também para arquivos grandes e chamadas de API que retornam JSON, pois não é necessário carregar todo o conteúdo na memória antes de decodificar.
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var data models.Rates

	err = json.NewDecoder(file).Decode(&data)
	check(err)

	// GET ARGS
	args := os.Args[1:] // ignore the first argument (./program_name)

	const MIN_ARGS_LENGTH = 2

	if len(args) < MIN_ARGS_LENGTH {
		fmt.Println("Informe: [valor_em_reais] [moeda_destino]. Ex: ./converter 100 USD")
		return
		// os.Exit(1) // Close the program with an error code, and don't execute the defer statements
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Valor inválido. Informe um número válido.")
		return
	}

	// GET CURRENCY
	currency := args[1]

	rate, ok := data.Rates[currency]
	if !ok {
		fmt.Printf("Moeda inválida. Informe uma das seguintes moedas: %v\n", getCurrencies(data.Rates))
		return
	}

	fmt.Println("Amount:", amount)
	fmt.Println("Currency:", currency)
	fmt.Println("Rate:", rate)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getCurrencies(rates map[string]float64) []string {
	currencies := make([]string, 0, len(rates))

	for currency := range rates {
		currencies = append(currencies, currency)
	}

	return currencies
}

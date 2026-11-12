package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

	fmt.Println(data.Base)
	fmt.Println(data.Date)
	fmt.Println(data.Rates)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

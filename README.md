# Conversor de Moedas

Aplicação de linha de comando escrita em Go para converter valores em reais (BRL) para outras moedas usando taxas de câmbio armazenadas localmente em JSON.

## Funcionalidades

- Converte valores em BRL para moedas disponíveis no arquivo de `rates.json`(taxas).
- Aceita o código da moeda em letras maiúsculas ou minúsculas.
- Exibe uma mensagem de uso quando os argumentos são insuficientes.
- Lista as moedas disponíveis quando o código informado é inválido.
- Mantém as taxas em um arquivo simples e fácil de atualizar.

## Pré-requisitos

- [Go](https://go.dev/dl/) 1.27 ou superior.

## Instalação

Clone o repositório e entre no diretório do projeto:

```bash
git clone https://github.com/PabloXT14/go-currency-converter.git
cd go-currency-converter
```

Baixe as dependências e valide o projeto:

```bash
go mod download
go test ./...
```

## Uso

Execute diretamente com `go run`:

```bash
go run . 100 USD
```

Saída:

```text
100.00 BRL = 15.10 USD
```

Para gerar um executável:

```bash
go build -o converter .
./converter 100 EUR
```

O formato do comando é:

```text
./converter [valor_em_reais] [moeda_destino]
```

Exemplos:

```bash
go run . 250 eur
go run . 1000 JPY
```

Se a moeda não existir, o programa informa os códigos disponíveis:

```text
moeda inválida "XYZ". Moedas disponíveis: ...
```

## Taxas de câmbio

As taxas utilizadas pela aplicação ficam em
[`data/rates.json`](data/rates.json). O campo `base` define a moeda de origem
(atualmente `BRL`), `date` registra a data das taxas e `rates` contém o fator
de conversão de cada moeda.

Para atualizar as taxas, edite esse arquivo mantendo a estrutura esperada:

```json
{
  "base": "BRL",
  "date": "AAAA-MM-DD",
  "rates": {
    "USD": 0.151
  }
}
```

Os valores são estáticos: o programa não consulta uma API externa durante a
execução.

## Estrutura do projeto

```text
.
├── data/
│   └── rates.json       # Taxas de câmbio
├── models/
│   └── rates.go        # Modelo dos dados de taxas
├── main.go             # Entrada da aplicação e lógica do CLI
├── go.mod              # Configuração do módulo Go
└── README.md
```

## Desenvolvimento

Depois de modificar o código ou as taxas, execute:

```bash
go test ./...
go vet ./...
```

Contribuições são bem-vindas. Abra uma issue para discutir uma mudança maior
ou envie um pull request com uma descrição clara do que foi alterado.

## Licença

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](./LICENSE) para mais informações.

<p align="center">
  Feito com 💜 por Pablo Alan 👋🏽 <a href="https://www.linkedin.com/in/pabloalan/" target="_blank">Entre em contato!</a>
</p>
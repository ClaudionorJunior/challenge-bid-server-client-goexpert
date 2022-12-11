# Descrição do projeto
Esse projeto faz parte do primeiro desafio do curso Go Expert da FullCycle.

# Os requisitos para cumprir este desafio são:
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: `https://economia.awesomeapi.com.br/json/last/USD-BRL` e em seguida deverá retornar no formato JSON o resultado para o cliente.

O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context".
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: `/cotacao` e a porta a ser utilizada pelo servidor HTTP será a 8080.

# Rodando o projeto
`cd server` e `go run main.go`
`cd client` e `go run main.go`

## Autor

![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)
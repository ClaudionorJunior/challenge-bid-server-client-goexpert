# Descrição do projeto
Esse projeto faz parte do primeiro desafio do curso Go Expert da [FullCycle](https://fullcycle.com.br/).<br>
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/ClaudionorJunior/challenge-bid-server-client-goexpert">

# Os requisitos para cumprir este desafio são:
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: `https://economia.awesomeapi.com.br/json/last/USD-BRL` e em seguida deverá retornar no formato JSON o resultado para o cliente.

O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context".
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: `/cotacao` e a porta a ser utilizada pelo servidor HTTP será a 8080.

# Rodando o projeto
```sh
cd server && go run main.go
```
```sh
cd client && go run main.go
```

# Autor
<view style="display:flex;">
  <view style="display:flex; flex-direction:column; align-items:center;">
    <img src="https://avatars.githubusercontent.com/u/82416762?v=4" width=60 style="border-radius: 30px"><br>
    <a href="https://github.com/ClaudionorJunior">Github</a>
    <a href="https://www.linkedin.com/in/claudionorsilva">Linkedin</a>
  </view>
</view>
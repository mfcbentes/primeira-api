# API RESTful em Go com Gorilla Mux

Este projeto é uma API RESTful simples escrita em Go usando o pacote Gorilla Mux. Ele fornece operações CRUD (Create, Read, Update, Delete) para gerenciar informações de pessoas, onde cada pessoa tem um endereço associado.

## Estruturas de Dados

O projeto define duas estruturas de dados, `Person` e `Address`.

```go
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Person struct {
	ID        int     `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Address   Address `json:"address"`
}
```

## Endpoints

A API define os seguintes endpoints:

- `GET /people`: Retorna uma lista de todas as pessoas.
- `GET /people/{id}`: Retorna a pessoa com o ID especificado.
- `POST /people`: Cria uma nova pessoa com os dados fornecidos no corpo da solicitação.
- `DELETE /people/{id}`: Exclui a pessoa com o ID especificado.

## Executando o Projeto

Para executar o projeto, você precisa ter Go instalado em sua máquina. Depois de clonar o repositório, você pode iniciar o servidor com o seguinte comando:

```bash
go run main.go
```

O servidor começará a ouvir na porta 8000. Você pode interagir com a API usando uma ferramenta como curl ou Postman.

## Nota

Este é um projeto de exemplo e os dados são armazenados na memória, portanto, serão perdidos quando o servidor for desligado. Para um aplicativo real, você provavelmente gostaria de usar um banco de dados para armazenar os dados.

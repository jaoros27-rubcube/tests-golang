package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	name  string
	email string
	cpf   string
}

func (c Client) display() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.name)
}

type InternationalClient struct {
	client  Client
	country string
}

func main() {
	client := Client{
		name:  "João Rós",
		email: "joao.guilherme@rubcube.com",
		cpf:   "12345678909",
	}
	fmt.Println(client)

	client2 := Client{"João Ros Garrido", "joao.vitor@rubcube.com", "98765432101"}
	fmt.Printf("Nome: %s. Email: %s. CPF: %s\n", client2.name, client2.email, client2.cpf)

	client3 := InternationalClient{
		client: Client{
			name:  "Jean Ros",
			email: "jean.junior@rubcube.com",
			cpf:   "12345678909",
		},
		country: "EUA",
	}
	fmt.Printf("Nome: %s. Email: %s. CPF: %s. País: %s\n", client3.client.name, client3.client.email, client3.client.cpf, client3.country)

	client.display()
	client3.client.display()

	clientJson, err := json.Marshal(client3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clientJson))

	jsonclient4 := `{"name": "Jean Ros", "email": "jean.junior@rubcube.com", "cpf": "12345678909", "pais": "EUA"}`
	client4 := InternationalClient{}

	json.Unmarshal([]byte(jsonclient4), &client4)
	fmt.Println(client4)
}

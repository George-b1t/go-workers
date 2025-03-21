package main

import (
	"go-threads/internal/app/client"
	"log"
)

func main() {
	client := client.NewClient()

	// Conectar ao servidor master
	if err := client.Connect("localhost:12345"); err != nil {
		log.Fatal(err)
	}

	// Iniciar interação com o usuário
	client.Start()
}

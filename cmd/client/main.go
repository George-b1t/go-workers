package main

import (
	"go-threads/internal/app/client"
	"log"
)

// main inicializa e executa o cliente que se conecta ao servidor
func main() {
	client := client.NewClient()

	// Tenta se conectar ao servidor na porta 12345
	if err := client.Connect("localhost:12345"); err != nil {
		log.Fatal(err)
	}

	// Inicia a comunicação com o servidor
	client.Start()
}

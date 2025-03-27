package main

import (
	"go-threads/internal/app/server"
	"log"
)

// main inicializa um servidor básico com buffer de tarefas
func main() {
	// Cria um novo servidor com buffer para 100 tarefas
	server := server.NewServer(100)

	// Inicia o servidor na porta 12345
	if err := server.Start(":12345"); err != nil {
		log.Fatal(err)
	}
}

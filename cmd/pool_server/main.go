package main

import (
	server "go-threads/internal/app/pool_server"
	"log"
)

// main inicializa o servidor com uma pool de workers
func main() {
	// Cria uma pool com 3 workers e buffer de 10 tarefas
	srvPool := server.NewServerPool(3, 10)

	// Inicia o servidor na porta 12345
	err := srvPool.Start(":12345")
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor (pool): %v", err)
	}
}

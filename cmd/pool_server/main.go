package main

import (
	server "go-threads/internal/app/pool_server"
	"log"
)

func main() {
	srvPool := server.NewServerPool(3, 10)

	err := srvPool.Start(":12345")
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor (pool): %v", err)
	}
}

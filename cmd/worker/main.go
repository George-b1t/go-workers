package main

import (
	"log"
	"net"
	"go-threads/internal/app/worker"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("Erro ao conectar ao master:", err)
	}
	defer conn.Close()

	log.Println("Conectado ao master como worker")
	worker := worker.NewWorker(conn)
	if err := worker.Start(); err != nil {
		log.Fatal(err)
	}
}

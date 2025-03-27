package main

import (
	"go-threads/internal/app/worker"
	"log"
	"net"
)

// main inicializa um worker que se conecta ao servidor master
func main() {
	// Estabelece conexão com o servidor master na porta 12345
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("Erro ao conectar ao master:", err)
	}
	defer conn.Close()

	log.Println("Conectado ao master como worker")

	// Inicializa e inicia o worker para processar tarefas recebidas
	worker := worker.NewWorker(conn)
	if err := worker.Start(); err != nil {
		log.Fatal(err)
	}
}

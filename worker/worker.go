package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Conecta ao servidor master.
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("Erro ao conectar ao master:", err)
	}
	defer conn.Close()
	log.Println("Conectado ao master como worker")

	// Envia a identificação.
	fmt.Fprintf(conn, "worker\n")

	reader := bufio.NewReader(conn)
	for {
		// Aguarda a chegada de uma tarefa.
		task, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Erro ao ler tarefa:", err)
			os.Exit(1)
		}
		task = strings.TrimSpace(task)
		log.Printf("Tarefa recebida: %s\n", task)
		// Simula o processamento (tempo aleatório entre 1 e 3 segundos)
		processingTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(processingTime)
		// Simula falha com 20% de chance.
		if rand.Float32() < 0.2 {
			log.Println("Falha ao processar tarefa")
			fmt.Fprintf(conn, "fail\n")
		} else {
			result := fmt.Sprintf("resultado do '%s' processado", task)
			log.Printf("Tarefa concluída: %s\n", result)
			fmt.Fprintf(conn, "%s\n", result)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Conecta ao servidor master.
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("Erro ao conectar ao master:", err)
	}
	defer conn.Close()
	log.Println("Conectado ao master como cliente")

	// Envia a identificação.
	fmt.Fprintf(conn, "client\n")

	readerServer := bufio.NewReader(conn)
	readerInput := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite a tarefa a ser processada (ou 'bye' para sair): ")
		task, err := readerInput.ReadString('\n')
		if err != nil {
			log.Fatal("Erro ao ler entrada do usuário:", err)
		}
		task = strings.TrimSpace(task)
		// Envia a tarefa para o master.
		_, err = fmt.Fprintf(conn, "%s\n", task)
		if err != nil {
			log.Fatal("Erro ao enviar tarefa para o master:", err)
		}
		// Lê a resposta do servidor.
		result, err := readerServer.ReadString('\n')
		if err != nil {
			log.Fatal("Erro ao ler resposta do master:", err)
		}
		result = strings.TrimSpace(result)
		fmt.Printf("Resultado recebido: %s\n", result)
		// Se o usuário digitou "bye", encerra a conexão.
		if strings.ToLower(task) == "bye" {
			break
		}
	}
	fmt.Println("Cliente encerrado.")
}

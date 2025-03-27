package client

import (
	"bufio"
	"fmt"
	"go-threads/internal/app/worker"
	"log"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn         net.Conn
	serverReader *bufio.Reader
	inputReader  *bufio.Reader
}

func NewClient() *Client {
	return &Client{
		inputReader: bufio.NewReader(os.Stdin),
	}
}

func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao servidor: %w", err)
	}

	c.conn = conn
	c.serverReader = bufio.NewReader(conn)

	if _, err := fmt.Fprintln(conn, "client"); err != nil {
		conn.Close()
		return fmt.Errorf("erro ao se identificar como cliente: %w", err)
	}

	return nil
}

func (c *Client) Start() {
	defer c.conn.Close()
	log.Println("Conectado ao servidor master")

	for {
		fmt.Println("Digite a tarefa a ser processada + variavel (ou 'bye' para sair): ")
		fmt.Println("Exemplo: uppercase:ola")
		for key := range worker.Tasks {
			fmt.Println("-", key)
		}
		task, err := c.inputReader.ReadString('\n')
		if err != nil {
			log.Fatalf("Erro ao ler entrada: %v", err)
		}

		parts := strings.Split(task, ":")

		if len(parts) != 2 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Formato inválido")
			continue
		}

		tarefa := parts[0]

		if _, ok := worker.Tasks[tarefa]; !ok {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Tarefa não encontrada")
			continue
		}

		task = strings.TrimSpace(task)
		if err := c.SendTask(task); err != nil {
			log.Fatal(err)
		}

		result, err := c.ReceiveResult()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("\033[H\033[2J")

		fmt.Printf("Resultado recebido: %s\n", result)

		if strings.EqualFold(task, "bye") {
			break
		}
	}

	fmt.Println("Conexão encerrada")
}

func (c *Client) SendTask(task string) error {
	if _, err := fmt.Fprintln(c.conn, task); err != nil {
		return fmt.Errorf("erro ao enviar tarefa: %w", err)
	}
	return nil
}

func (c *Client) ReceiveResult() (string, error) {
	result, err := c.serverReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %w", err)
	}
	return strings.TrimSpace(result), nil
}

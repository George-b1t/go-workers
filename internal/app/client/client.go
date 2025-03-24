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

// Client representa um cliente conectado ao servidor master.
type Client struct {
	conn         net.Conn
	serverReader *bufio.Reader
	inputReader  *bufio.Reader
}

// NewClient cria uma nova instância de Client.
func NewClient() *Client {
	return &Client{
		inputReader: bufio.NewReader(os.Stdin),
	}
}

// Connect estabelece conexão com o servidor master e se identifica como cliente.
func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao servidor: %w", err)
	}
	
	c.conn = conn
	c.serverReader = bufio.NewReader(conn)
	
	// Envia identificação como cliente
	if _, err := fmt.Fprintln(conn, "client"); err != nil {
		conn.Close()
		return fmt.Errorf("erro ao se identificar como cliente: %w", err)
	}
	
	return nil
}

// Start inicia o loop de interação com o usuário.
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
		
		task = strings.TrimSpace(task)
		if err := c.sendTask(task); err != nil {
			log.Fatal(err)
		}
		
		result, err := c.receiveResult()
		if err != nil {
			log.Fatal(err)
		}

		//limpa o terminal
		fmt.Print("\033[H\033[2J")
		
		fmt.Printf("Resultado recebido: %s\n", result)
		
		if strings.EqualFold(task, "bye") {
			break
		}
	}
	
	fmt.Println("Conexão encerrada")
}

// sendTask envia uma tarefa para o servidor.
func (c *Client) sendTask(task string) error {
	if _, err := fmt.Fprintln(c.conn, task); err != nil {
		return fmt.Errorf("erro ao enviar tarefa: %w", err)
	}
	return nil
}

// receiveResult recebe o resultado do processamento do servidor.
func (c *Client) receiveResult() (string, error) {
	result, err := c.serverReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %w", err)
	}
	return strings.TrimSpace(result), nil
}
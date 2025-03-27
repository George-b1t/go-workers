package poolserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Task representa uma tarefa com um canal para o resultado
type Task struct {
	Job        string
	ResultChan chan string
}

// ServerPool gerencia a pool de workers e tarefas
type ServerPool struct {
	tasks chan Task
}

// NewServerPool cria uma nova pool de workers com tamanho específico
func NewServerPool(poolSize, bufferSize int) *ServerPool {
	sp := &ServerPool{
		tasks: make(chan Task, bufferSize),
	}

	for i := 0; i < poolSize; i++ {
		go sp.worker(i)
	}

	return sp
}

// worker é a função executada por cada worker que processa tarefas
func (sp *ServerPool) worker(workerID int) {
	for t := range sp.tasks {
		result, err := ProcessTask(t.Job)
		if err != nil {
			t.ResultChan <- "fail"
		} else {
			t.ResultChan <- result
		}
	}
}

// Start inicia o servidor e escuta conexões na porta informada
func (sp *ServerPool) Start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("erro ao iniciar o servidor: %w", err)
	}
	defer listener.Close()

	log.Println("Servidor (goroutine pool) iniciado na porta", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go sp.handleClient(conn)
	}
}

// handleClient gerencia as interações com cada cliente conectado
func (sp *ServerPool) handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	log.Printf("Cliente conectado: %s\n", conn.RemoteAddr().String())

	idLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Erro ao ler identificação: %v\n", err)
		return
	}
	log.Printf("Identificação recebida: %s\n", strings.TrimSpace(idLine))

	for {
		taskMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Erro ao ler tarefa do cliente: %v\n", err)
			return
		}

		taskMsg = strings.TrimSpace(taskMsg)

		if strings.EqualFold(taskMsg, "bye") {
			conn.Write([]byte("bye\n"))
			log.Printf("Cliente %s desconectou.\n", conn.RemoteAddr().String())
			return
		}

		// Cria uma tarefa e espera pelo resultado do worker
		resultChan := make(chan string)
		sp.tasks <- Task{
			Job:        taskMsg,
			ResultChan: resultChan,
		}

		result := <-resultChan
		_, err = conn.Write([]byte(result + "\n"))
		if err != nil {
			log.Printf("Erro ao enviar resposta ao cliente: %v\n", err)
			return
		}
		log.Printf("Resultado enviado ao cliente %s: %s\n", conn.RemoteAddr().String(), result)
	}
}

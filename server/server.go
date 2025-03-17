package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Worker representa um worker registrado com sua conexão.
type Worker struct {
	conn net.Conn
}

// Server gerencia conexões de workers e clientes.
type Server struct {
	workerPool chan *Worker
	mu         sync.Mutex
}

// NewServer cria uma instância do Server com um pool de workers.
func NewServer(poolSize int) *Server {
	return &Server{
		workerPool: make(chan *Worker, poolSize),
	}
}

// Start inicia o servidor na porta especificada.
func (s *Server) Start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("erro ao iniciar o servidor: %w", err)
	}
	defer listener.Close()

	log.Println("Servidor iniciado na porta", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection identifica se a conexão é de um worker ou cliente.
func (s *Server) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	idMsg, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Erro ao ler identificação:", err)
		conn.Close()
		return
	}
	idMsg = strings.TrimSpace(idMsg)

	if idMsg == "worker" {
		s.registerWorker(conn)
	} else if idMsg == "client" {
		s.handleClient(conn)
	} else {
		log.Println("Identificação desconhecida:", idMsg)
		conn.Close()
	}
}

// registerWorker adiciona um worker ao pool e mantém sua conexão aberta.
func (s *Server) registerWorker(conn net.Conn) {
	worker := &Worker{conn: conn}
	log.Printf("Worker registrado: %s\n", conn.RemoteAddr().String())

	s.mu.Lock()
	s.workerPool <- worker
	s.mu.Unlock()

	select {} // Mantém a conexão ativa indefinidamente
}

// handleClient processa as tarefas enviadas por um cliente.
func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	log.Printf("Cliente conectado: %s\n", conn.RemoteAddr().String())

	for {
		taskMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Erro ao ler tarefa do cliente %s: %v\n", conn.RemoteAddr().String(), err)
			return
		}

		taskMsg = strings.TrimSpace(taskMsg)
		if strings.ToLower(taskMsg) == "bye" {
			conn.Write([]byte("bye\n"))
			log.Printf("Cliente %s desconectou.\n", conn.RemoteAddr().String())
			return
		}

		log.Printf("Tarefa recebida do cliente %s: %s\n", conn.RemoteAddr().String(), taskMsg)
		result := s.assignTask(taskMsg)
		_, err = conn.Write([]byte(result + "\n"))
		if err != nil {
			log.Printf("Erro ao enviar resposta para o cliente %s: %v\n", conn.RemoteAddr().String(), err)
			return
		}
		log.Printf("Resultado enviado ao cliente %s: %s\n", conn.RemoteAddr().String(), result)
	}
}

// assignTask atribui a tarefa a um worker disponível e trata falhas.
func (s *Server) assignTask(task string) string {
	for {
		s.mu.Lock()
		if len(s.workerPool) == 0 {
			s.mu.Unlock()
			log.Println("Nenhum worker disponível. Aguardando...")
			time.Sleep(2 * time.Second)
			continue
		}
		worker := <-s.workerPool
		s.mu.Unlock()

		log.Printf("Atribuindo tarefa '%s' ao worker %s\n", task, worker.conn.RemoteAddr().String())
		_, err := worker.conn.Write([]byte(task + "\n"))
		if err != nil {
			log.Println("Erro ao enviar tarefa para o worker:", err)
			continue
		}

		worker.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		responseReader := bufio.NewReader(worker.conn)
		resp, err := responseReader.ReadString('\n')
		if err != nil {
			log.Println("Erro ao ler resposta do worker:", err)
			continue
		}
		resp = strings.TrimSpace(resp)
		if resp == "fail" {
			log.Printf("Worker %s falhou na tarefa '%s'. Reatribuindo...\n", worker.conn.RemoteAddr().String(), task)
			s.mu.Lock()
			s.workerPool <- worker
			s.mu.Unlock()
			continue
		}

		s.mu.Lock()
		s.workerPool <- worker
		s.mu.Unlock()
		return resp
	}
}

func main() {
	server := NewServer(100)
	if err := server.Start(":12345"); err != nil {
		log.Fatal(err)
	}
}

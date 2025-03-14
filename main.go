package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

// Worker representa um worker registrado com sua conexão.
type Worker struct {
	conn net.Conn
}

var workerPool = make(chan *Worker, 100) // Pool de workers disponíveis

// handleConnection identifica se a conexão é de um worker ou cliente e age de acordo.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Lê a identificação inicial
	idMsg, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Erro ao ler identificação:", err)
		return
	}
	idMsg = strings.TrimSpace(idMsg)

	if idMsg == "worker" {
		// Registra o worker e o coloca no pool.
		log.Printf("Worker registrado: %s\n", conn.RemoteAddr().String())
		worker := &Worker{conn: conn}
		workerPool <- worker
		// Mantém a conexão aberta indefinidamente para que o worker possa receber tarefas.
		select {}
	} else if idMsg == "client" {
		log.Printf("Cliente conectado: %s\n", conn.RemoteAddr().String())
		// Loop para processar várias tarefas do mesmo cliente
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
			// Processa a tarefa e envia o resultado
			result := assignTask(taskMsg)
			_, err = conn.Write([]byte(result + "\n"))
			if err != nil {
				log.Printf("Erro ao enviar resposta para o cliente %s: %v\n", conn.RemoteAddr().String(), err)
				return
			}
			log.Printf("Resultado enviado ao cliente %s: %s\n", conn.RemoteAddr().String(), result)
		}
	} else {
		log.Println("Identificação desconhecida:", idMsg)
	}
}

// assignTask atribui uma tarefa a um worker disponível. Se o worker falhar (responder "fail"),
// a tarefa é reatribuída a outro worker.
func assignTask(task string) string {
	for {
		worker := <-workerPool // Aguarda por um worker disponível
		log.Printf("Atribuindo tarefa '%s' ao worker %s\n", task, worker.conn.RemoteAddr().String())
		_, err := worker.conn.Write([]byte(task + "\n"))
		if err != nil {
			log.Println("Erro ao enviar tarefa para o worker:", err)
			// Descartamos esse worker e tentamos com o próximo
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
			workerPool <- worker // Retorna o worker ao pool mesmo em caso de falha
			continue
		} else {
			workerPool <- worker // Retorna o worker para o pool após uso
			return resp
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
	defer listener.Close()
	log.Println("Servidor master iniciado na porta 12345")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleConnection(conn)
	}
}

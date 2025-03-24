package worker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Worker representa um worker que processa tarefas recebidas do master.
type Worker struct {
	conn io.ReadWriter
}

// NewWorker cria uma nova instância de Worker.
func NewWorker(conn io.ReadWriter) *Worker {
	return &Worker{conn: conn}
}

// ProcessTask recebe uma tarefa, simula o processamento e retorna o resultado ou erro.
func ProcessTask(task string) (string, error) {
	task = strings.TrimSpace(task)
	if task == "" {
		return "", fmt.Errorf("tarefa vazia")
	}

	
	parts := strings.Split(task, ":")
	
	if len(parts) != 2 {
		fmt.Println("Formato inválido")
		return "", fmt.Errorf("formato inválido")
	}
	
	tarefa := parts[0]
	texto := parts[1]
	
	if _, ok := Tasks[tarefa]; !ok {
		fmt.Println("Tarefa não encontrada")
		return "", fmt.Errorf("tarefa não encontrada")
	}

	processingTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(processingTime)
	
	result, err := Tasks[tarefa](texto)
	if err != nil {
		return "", fmt.Errorf("falha ao processar tarefa: %w", err)
	}

	if rand.Float32() < 0.2 {
		return "", fmt.Errorf("falha no processamento")
	}

	return result, nil
}

// Start inicia o loop de recebimento e processamento de tarefas.
func (w *Worker) Start() error {
	// Envia a identificação do worker
	_, err := fmt.Fprintln(w.conn, "worker")
	if err != nil {
		return fmt.Errorf("erro ao enviar identificação: %w", err)
	}

	reader := bufio.NewReader(w.conn)
	for {
		task, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Conexão encerrada pelo servidor")
				return nil
			}
			return fmt.Errorf("erro ao ler tarefa: %w", err)
		}

		result, err := ProcessTask(task)

		if err != nil {
			log.Println("Falha ao processar tarefa:", err)
			fmt.Fprintln(w.conn, "fail")
		} else {
			log.Println("Tarefa concluída:", result)
			fmt.Fprintln(w.conn, result)
		}
	}
}

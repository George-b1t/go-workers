package poolserver

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ProcessTask processa uma tarefa específica e retorna o resultado
func ProcessTask(task string) (string, error) {
	task = strings.TrimSpace(task)
	if task == "" {
		return "", fmt.Errorf("tarefa vazia")
	}

	// Divide o comando e o texto da tarefa
	parts := strings.Split(task, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("formato inválido, use algo como uppercase:ola")
	}

	command := parts[0]
	text := parts[1]

	// Simula um processamento com duração aleatória entre 10 e 12 segundos
	time.Sleep(time.Duration(rand.Intn(3)+10) * time.Second)

	// Executa o comando solicitado
	switch command {
	case "uppercase":
		return strings.ToUpper(text), nil
	case "lowercase":
		return strings.ToLower(text), nil
	case "reverse":
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case "caesar":
		return caesarCipher(text), nil
	default:
		return "", fmt.Errorf("comando desconhecido: %s", command)
	}
}

// caesarCipher aplica uma cifra de César simples deslocando caracteres em 3 posições
func caesarCipher(text string) string {
	result := []rune(text)
	for i, char := range result {
		if char >= 'a' && char <= 'z' {
			result[i] = rune((int(char-'a')+3)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result[i] = rune((int(char-'A')+3)%26 + 'A')
		}
	}
	return string(result)
}

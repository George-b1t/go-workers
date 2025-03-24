package worker

import (
	"strings"
	"unicode"
)

// Mapa de funções públicas
var Tasks = map[string]func(string) (string, error){
	"reverse":   Reverse,
	"uppercase": Uppercase,
	"lowercase": Lowercase,
	"caesar":    caesarCipher,
}

// Função para inverter uma string
func Reverse(s string) (string, error) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

// Função para transformar em maiúsculas
func Uppercase(s string) (string, error) {
	return strings.ToUpper(s), nil
}

// Função para transformar em minúsculas
func Lowercase(s string) (string, error) {
	return strings.ToLower(s), nil
}

func caesarCipher(text string) (string, error) {
	result := []rune(text)

	for i, char := range result {
		if unicode.IsLetter(char) { // Verifica se é uma letra
			base := 'A'
			if unicode.IsLower(char) {
				base = 'a'
			}
			// Aplica o deslocamento, garantindo que permaneça dentro do alfabeto
			result[i] = rune((int(char-base)+3)%26 + int(base))
		}
	}

	return string(result), nil
}

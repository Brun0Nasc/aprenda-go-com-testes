package main

import "fmt"

const (
	espanhol = "espanhol"
	frances = "francês"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol = "Hola, "
	prefixoFrances = "Bonjour, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == frances {
		return prefixoFrances + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo", "espanhol"))
}

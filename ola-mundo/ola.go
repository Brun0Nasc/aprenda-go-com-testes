package main

import "fmt"

const (
	espanhol            = "espanhol"
	frances             = "francês"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoFrances      = "Bonjour, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case frances:
		prefixo = prefixoFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	}

	return prefixo + nome
}

func main() {
	fmt.Println(Ola("mundo", "espanhol"))
}

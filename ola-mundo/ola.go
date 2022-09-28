package main

import "fmt"

const (
	espanhol            = "espanhol"
	frances             = "francês"
	russo               = "russo"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoFrances      = "Bonjour, "
	prefixoRusso        = "Привет, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case russo:
		prefixo = prefixoRusso
	default:
		prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", "espanhol"))
}

package main

import "fmt"

const prefixoOla = "Olá, "

func Ola(nome string) string{
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoOla + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}

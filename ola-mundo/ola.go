package main

import "fmt"

const prefixoOla = "Olá, "

func Ola(nome string) string{
	return prefixoOla + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}

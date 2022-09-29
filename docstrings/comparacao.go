package docstrings

import "strings"

// Comparar vai chamar a função Compare do pacote strings e retornar um valor inteiro que indica a distância entre as letras inseridas
func Comparar(x, y string) int {
	return strings.Compare(x, y)
}
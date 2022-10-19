package main

import (
	"fmt"
	"io"
	"os"
)

const (
	ultimaPalavra = "Vai!"
	inicioContagem = 3
)

func Contagem(saida io.Writer) { //* io.Writer é a interface que vai receber uma fatia de bytes que vai ser mandada para qualquer saída especificada
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	Contagem(os.Stdout) //* os.Stdout vai definir que a saída desse resultado vai ser no Terminal
}

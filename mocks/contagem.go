package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct {}

const (
	ultimaPalavra  = "Vai!"
	inicioContagem = 3
)

func Contagem(saida io.Writer, sleeper Sleeper) { //* io.Writer é a interface que vai receber uma fatia de bytes que vai ser mandada para qualquer saída especificada
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper) //* os.Stdout vai definir que a saída desse resultado vai ser no Terminal
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

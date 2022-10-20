package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pausa()
}

type SleeperPadrao struct {}

const (
	ultimaPalavra  = "Vai!"
	inicioContagem = 3
)

func Contagem(saida io.Writer, sleeper Sleeper) { //* io.Writer é a interface que vai receber uma fatia de bytes que vai ser mandada para qualquer saída especificada
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper) //* os.Stdout vai definir que a saída desse resultado vai ser no Terminal
}

func (d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}

package mocks

import (
	"bytes"
	"fmt"
)

func Contagem(saida *bytes.Buffer){
	fmt.Fprintf(saida, "3")
}

package mocks

import (
	"fmt"
	"io"
	"os"
)

func Contagem(saida io.Writer){
	fmt.Fprintf(saida, "3")
}

func main(){
	Contagem(os.Stdout)
}

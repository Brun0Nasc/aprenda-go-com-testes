package mocks

import (
	"fmt"
	"io"
	"os"
)

func Contagem(saida io.Writer){ //* io.Writer é a interface que vai receber uma fatia de bytes que vai ser mandada para qualquer saída especificada
	fmt.Fprintf(saida, "3")
}

func main(){
	Contagem(os.Stdout) //* os.Stdout vai definir que a saída desse resultado vai ser no Terminal
}

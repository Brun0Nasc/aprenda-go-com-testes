package inteiros

import (
	"fmt"
	"testing"
)

func ExampleAdiciona() {
	soma := Adiciona(5, 2)
	fmt.Println(soma)
	// Output: 7
}
func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

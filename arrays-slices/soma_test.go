package arraysslices

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("Resultado '%d', esperado '%d', dados '%v'", resultado, esperado, numeros)
		}
	})

}

func TestSomaTodoOResto(t *testing.T) {
	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{2, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado '%v', esperado '%v'", resultado, esperado)
	}
}

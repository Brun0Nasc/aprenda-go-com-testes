package arraysslices

import "testing"

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

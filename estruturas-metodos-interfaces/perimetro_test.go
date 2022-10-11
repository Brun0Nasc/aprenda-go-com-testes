package estruturasmetodosinterfaces

import "testing"

func TestAreaPerimetro(t *testing.T) {
	verificaResultado := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	}

	t.Run("verifica perimetro", func(t *testing.T) {
		resultado := Perimetro(10.0, 10.0)
		esperado := 40.0

		verificaResultado(t, resultado, esperado)
	})

	t.Run("verifica area", func(t *testing.T) {
		resultado := Area(12.0, 6.0)
		esperado := 72.0

		verificaResultado(t, resultado, esperado)
	})
}

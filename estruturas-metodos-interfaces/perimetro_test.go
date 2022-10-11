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
		retangulo := Retangulo{10.0, 10.0}
		resultado := retangulo.Perimetro()
		esperado := 40.0

		verificaResultado(t, resultado, esperado)
	})

	t.Run("verifica area", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0

		verificaResultado(t, resultado, esperado)
	})

	t.Run("verifica area de circulos", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := 314.1592653589793

		verificaResultado(t, resultado, esperado)
	})
}

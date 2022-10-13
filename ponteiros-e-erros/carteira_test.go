package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	verificaTeste := func (t *testing.T, resultado, esperado Bitcoin)  {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))

		resultado := carteira.Saldo()

		esperado := Bitcoin(10)

		verificaTeste(t, resultado, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))

		resultado := carteira.Saldo()

		esperado := Bitcoin(10)

		verificaTeste(t, resultado, esperado)
	})

}

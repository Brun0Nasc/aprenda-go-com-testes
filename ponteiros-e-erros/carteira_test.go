package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	//* método auxiliar de teste para verificar erros
	confirmaErro := func(t *testing.T, valor error, esperado string) {
		t.Helper()
		if valor == nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu.")
		}

		resultado := valor.Error()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	//* método auxiliar de teste para verificar resultados do teste
	verificaTeste := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		if carteira.saldo != esperado {
			t.Errorf("resultado %s, esperado %s", carteira.saldo, esperado)
		}
	}

	//* testes propriamente ditos
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))

		esperado := Bitcoin(10)

		verificaTeste(t, carteira, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))

		esperado := Bitcoin(10)

		verificaTeste(t, carteira, esperado)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		verificaTeste(t, carteira, saldoInicial)
		confirmaErro(t, erro, "saldo insuficiente para retirada")
	})

}

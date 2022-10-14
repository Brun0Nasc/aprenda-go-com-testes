package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	//* testes propriamente ditos

	//! Testando função de depositar
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		verificaTeste(t, carteira, Bitcoin(10))
	})

	//! Testando função de retirar com saldo suficiente
	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))

		verificaTeste(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	//! Testando função de retirar com saldo insuficiente
	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		verificaTeste(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})

}

// * método auxiliar de teste para verificar resultados do teste
func verificaTeste(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	if carteira.saldo != esperado {
		t.Errorf("resultado %s, esperado %s", carteira.saldo, esperado)
	}
}

// * para testar quando não tem erro
func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

// * método auxiliar de teste para verificar erros
func confirmaErro(t *testing.T, resultado error, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu.")
	}

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

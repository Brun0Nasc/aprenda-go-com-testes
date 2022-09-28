package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Bruno", "francês")
		esperado := "Bonjour, Bruno"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em russo", func(t *testing.T) {
		resultado := Ola("Smirnoff", "russo")
		esperado := "Привет, Smirnoff"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}

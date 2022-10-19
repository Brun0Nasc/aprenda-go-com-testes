package mocks

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{} //* Um Buffer é um buffer de bytes de tamanho variável com métodos Read e Write. O valor zero para Buffer é um buffer vazio pronto para uso.

	Contagem(buffer)

	resultado := buffer.String()
	esperado := "3"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
package reflexao

import "testing"

func TestPercorre(t *testing.T) {

	esperado := "Chris"
	var resultado []string

	x := struct {
		Nome string
	}{esperado}

	percorre(x, func(entrada string) {
		resultado = append(resultado, entrada)
	})

	if len(resultado) != 1 {
		t.Errorf("número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
	}

	if resultado[0] != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado[0], esperado)
	}
}
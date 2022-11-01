package selec

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {

	servidorLento := criarServidorComAtraso(20 * time.Millisecond)
	servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

	defer servidorRapido.Close()
	defer servidorLento.Close()

	//* Ao chamar uma função com o prefixo defer, ela será chamada após o término da função que a contém.

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
	//* httptest.NewServer recebe um http.HandlerFunc que vamos enviar para uma função anônima.
	//* http.HandlerFunc é um tipo que se parece com isso: type HandlerFunc func(ResponseWriter, *Requisicao)
}

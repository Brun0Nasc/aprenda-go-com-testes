package selec

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {

	t.Run("compara a velocidade de servidores, retornando o endereço mais rápido", func(t *testing.T) {
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)
		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		defer servidorRapido.Close()
		defer servidorLento.Close()

		//* Ao chamar uma função com o prefixo defer, ela será chamada após o término da função que a contém.

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, err := Corredor(URLLenta, URLRapida)

		if err != nil {
			t.Fatalf("não esperava um erro, mas obteve um %v", err)
		}

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("retorna erro se o servidor não responder em 10s", func(t *testing.T) {
		servidor := criarServidorComAtraso(25 * time.Millisecond)

		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("esperava um erro mas não obtive um")
		}
	})

}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
	//* httptest.NewServer recebe um http.HandlerFunc que vamos enviar para uma função anônima.
	//* http.HandlerFunc é um tipo que se parece com isso: type HandlerFunc func(ResponseWriter, *Requisicao)
}

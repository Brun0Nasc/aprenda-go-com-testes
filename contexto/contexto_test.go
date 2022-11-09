package contexto

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct { //* mock spystore
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("não implementado")
}

func (s *SpyResponseWriter) WriteHeader(StatusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store foi cancelado")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	/*
	* Estamos simulando um processo lento onde construímos o resultado lentamente adicionando
	* a string, caractere por caractere em uma goroutine. Quando a goroutine termina seu
	* trabalho, ela escreve a string no channel data. A goroutine escuta o ctx.Done e irá
	* parar o trabalho se um sinal for enviado nesse channel.
	 */

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store não foi avisada para cancelar")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("store foi avisada para cancelar")
// 	}
// }

func TestHandler(t *testing.T) {
	data := "olá, mundo"

	t.Run("retorna os dados da store", func(t *testing.T) {
		store := SpyStore{response: data, t: t}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
		}

		// store.assertWasNotCancelled()
	})

	t.Run("avisa a store para cancelar o trabalho se a requisição for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("uma resposta não deveria ter sido escrita")
		}
	})
}

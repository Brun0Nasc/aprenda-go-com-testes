package contexto

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string //* Capturar dados
	Cancel() //* Cancelar requisição
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch() //* executamos o Fetch em uma goroutine para escrever o resultado em um novo channel(data)
		}()

		select{ //* select vai efetivamente correr para os dois processos assíncronos
		case d := <- data:
			fmt.Fprint(w, d)
		case <-ctx.Done(): //* retorna um canal que recebe um sinal quando o context estiver "done"(finalizado) ou "cancelled"(cancelado)
			store.Cancel()
		}
	}
}
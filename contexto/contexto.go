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
		fmt.Fprint(w, store.Fetch()) //* print dos dados capturados da requisição
	}
}
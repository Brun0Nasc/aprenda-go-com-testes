package selec

import (
	"net/http"
	"time"
)

func Corredor(a, b string) (vencedor string) {
	duracaoA := medirTempoResposta(a)
	duracaoB := medirTempoResposta(b)

	if duracaoA < duracaoB {
		return a
	}

	return b
}

func medirTempoResposta(URL string) time.Duration {
	inicio := time.Now()
	http.Get(URL)
	return time.Since(inicio)
}

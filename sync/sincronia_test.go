package sync

import (
	"sync"
	"testing"
)

func NovoContador() *Contador {
	return &Contador{}
}
func TestContador(t *testing.T) {
	verificaContador := func(t *testing.T, resultado *Contador, esperado int) {
		t.Helper()
		if resultado.Valor() != esperado {
			t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
		}
	}

	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := NovoContador()
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, contador, 3)
	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := NovoContador()

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verificaContador(t, contador, contagemEsperada)

		/*
			*Um WaitGroup aguarda por uma coleção de goroutines terminar seu processamento. 
			*A goroutine principal faz a chamada para o Add definir o número de goroutines 
			*que serão esperadas. Então, cada uma das goroutines é executada e chama Done 
			*quando termina sua execução. Ao mesmo tempo, Wait pode ser usado para bloquear
			*a execução até que todas as goroutines tenham terminado.
		*/
	})
}

package sync

import "sync"

type Contador struct {
	mu    sync.Mutex
	valor int
}

/*
	*Um Mutex é uma trava de exclusão mútua. O valor zero de um Mutex é um Mutex destravado.
*/

func (c *Contador) Incrementa() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}

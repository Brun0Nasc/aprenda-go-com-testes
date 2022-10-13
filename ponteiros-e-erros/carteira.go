package ponteiroseerros

import "fmt"

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(valor Bitcoin) {
	c.saldo -= valor
}

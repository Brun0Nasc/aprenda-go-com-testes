package ponteiroseerros

type Carteira struct {
	saldo int
}

func (c *Carteira) Depositar(valor int) {
	c.saldo += valor
}

func (c *Carteira) Saldo() int {
	return c.saldo
}

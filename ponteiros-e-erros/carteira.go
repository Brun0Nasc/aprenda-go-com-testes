package ponteiroseerros

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface { //* interface definida no pacote fmt
	String() string
}

func (b Bitcoin) String() string { //* vai definir como seu tipo é impresso quando utilizado com operador de string em prints
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return errors.New("saldo insuficiente para retirada")
	}
	c.saldo -= valor
	return nil
}

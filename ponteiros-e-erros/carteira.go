package ponteiroseerros

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("saldo insuficiente para retirada")

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
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}

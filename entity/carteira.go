package entity

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (c *Carteira) Depositar(quantidade Bitcoin) {

	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return errors.New("eita")
	}
	c.saldo -= quantidade
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

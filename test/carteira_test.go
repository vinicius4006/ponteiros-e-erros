package test

import (
	"testing"

	"github.com/vinicius4006/ponteiros-e-erros/entity"
)

func TestCarteira(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira entity.Carteira, esperado entity.Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado.String(), esperado.String())
		}
	}
	confirmaErro := func(t *testing.T, erro error) {
		t.Helper()
		if erro == nil {
			t.Error("Esperava um erro mas nenhum ocorreu")
		}
	}
	t.Run("Depositar", func(t *testing.T) {
		carteira := entity.Carteira{}
		carteira.Depositar(entity.Bitcoin(10))
		confirmaSaldo(t, carteira, entity.Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := entity.Carteira{}
		carteira.Depositar(entity.Bitcoin(20))
		carteira.Retirar(10)

		confirmaSaldo(t, carteira, entity.Bitcoin(10))
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := entity.Bitcoin(20)
		carteira := entity.Carteira{}
		carteira.Depositar(saldoInicial)

		erro := carteira.Retirar(entity.Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		confirmaErro(t, erro)

	})

}

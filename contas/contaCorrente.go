package contas

import (
	"github.com/alura/alr-go-banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Cliente
	NumeroAgencia int
	NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {

	var msg string
	podeSacar := valor <= c.saldo && valor > 0

	if podeSacar {

		c.saldo -= valor
		msg = "Saque efetuado com sucesso."

	} else {

		msg = "Saldo insuficiente."

	}

	return msg, c.saldo

}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {

	var msg string

	if valor > 0 {

		c.saldo += valor
		msg = "Depósito efetuado com sucesso."

	} else {

		msg = "Valor inválido."

	}

	return msg, c.saldo

}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) (string, bool) {

	var msg string
	var success bool

	if valor > 0 && valor <= c.saldo {

		c.Sacar(valor)
		contaDestino.Depositar(valor)
		msg = "Transferência feita com sucesso."
		success = true

	} else {

		msg = "Saldo insuficiente."
		success = false

	}

	return msg, success

}

func (c *ContaCorrente) ObterSaldo() float64 {

	return c.saldo

}
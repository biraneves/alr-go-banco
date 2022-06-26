package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
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

func main() {

	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 200}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	fmt.Println(contaDaSilvia.Transferir(150, &contaDoGustavo))

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
	
}
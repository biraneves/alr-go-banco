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

func (c *ContaCorrente) Sacar(valor float64) string {

	var msg string
	podeSacar := valor <= c.saldo && valor > 0

	if podeSacar {

		c.saldo -= valor
		msg = "Saque efetuado com sucesso."

	} else {

		msg = "Saldo insuficiente."

	}

	return msg

}

func main() {

	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(300))
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(800))
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(-1000))
	fmt.Println(contaDaSilvia.saldo)

}
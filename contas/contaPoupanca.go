package contas

import (
	"github.com/alura/alr-go-banco/clientes"
)

type ContaPoupanca struct {
	Titular			clientes.Cliente
	NumeroAgencia	int
	NumeroConta		int
	Operacao		int
	saldo			float64
}
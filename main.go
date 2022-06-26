package main

import (
	"fmt"
	"github.com/alura/alr-go-banco/contas"
	"github.com/alura/alr-go-banco/clientes"
)

func main() {

	clienteBruno := clientes.Cliente {
		Nome: "Bruno",
		CPF: "123.456.789-10",
		Profissao: "Desenvolvedor",
	}

	clienteSilvia := clientes.Cliente {
		Nome: "Silvia",
		CPF: "123.789.456-11",
		Profissao: "PO",
	}

	contaDoBruno := contas.ContaCorrente {
		Titular: clienteBruno,
		NumeroAgencia: 123,
		NumeroConta: 123456,
	}

	contaDaSilvia := contas.ContaCorrente {
		Titular: clienteSilvia,
		NumeroAgencia: 123,
		NumeroConta: 135246,
	}

	contaDoBruno.Depositar(200)
	contaDaSilvia.Depositar(400)

	fmt.Println(contaDoBruno)
	fmt.Println(contaDaSilvia)
	
}
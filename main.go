package main

import (
	"fmt"
	"github.com/alura/alr-go-banco/contas"
	"github.com/alura/alr-go-banco/clientes"
)

func main() {

	contaDoBruno := contas.ContaCorrente{
		Titular: clientes.Cliente{
			Nome: "Bruno",
			CPF: "123.456.789-10",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 123,
		NumeroConta: 123456,
		Saldo: 100,
	}

	fmt.Println(contaDoBruno)
	
}
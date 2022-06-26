package main

import (
	"fmt"
	"github.com/alura/alr-go-banco/contas"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 200}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	fmt.Println(contaDaSilvia.Transferir(150, &contaDoGustavo))

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
	
}
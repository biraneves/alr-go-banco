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

func main() {

	contaDoGuilherme := ContaCorrente{
		titular: "Guilherme", 
		numeroAgencia: 589, 
		numeroConta: 123456, 
		saldo: 128.45,
	}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroAgencia = 333

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDaCris)

}
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

	conta1 := ContaCorrente{
		titular: "Ubirajara Neves", 
		numeroAgencia: 589, 
		numeroConta: 123456, 
		saldo: 128.45,
	}

	fmt.Println(conta1)

}
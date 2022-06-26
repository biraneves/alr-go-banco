package contas

type ContaCorrente struct {
	Titular string
	NumeroAgencia int
	NumeroConta int
	Saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {

	var msg string
	podeSacar := valor <= c.Saldo && valor > 0

	if podeSacar {

		c.Saldo -= valor
		msg = "Saque efetuado com sucesso."

	} else {

		msg = "Saldo insuficiente."

	}

	return msg, c.Saldo

}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {

	var msg string

	if valor > 0 {

		c.Saldo += valor
		msg = "Depósito efetuado com sucesso."

	} else {

		msg = "Valor inválido."

	}

	return msg, c.Saldo

}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) (string, bool) {

	var msg string
	var success bool

	if valor > 0 && valor <= c.Saldo {

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
package main

type Doller struct {
	amount int
}

func (d *Doller) times(multipulier int) Doller {
	return Doller{d.amount * multipulier}
}

// 【Check】TDD本では引数をObject型にしていた
func (d *Doller) equals(doller Doller) bool {
	return d.amount == doller.amount
}

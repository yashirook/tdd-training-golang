package main

type Doller struct {
	amount int
}

func (d *Doller) times(multipulier int) Doller {
	return Doller{d.amount * multipulier}
}

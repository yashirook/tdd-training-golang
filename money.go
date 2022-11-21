package main

type Doller struct {
	amount int
}

func (d *Doller) times(multipulier int) {
	d.amount = multipulier * d.amount
}

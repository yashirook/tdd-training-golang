package main

type Doller struct {
	amount int
}

func (d *Doller) times(multipulier int) Doller {
	amount := d.amount * multipulier

	return Doller{amount}
}

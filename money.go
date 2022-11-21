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

type Franc struct {
	amount int
}

func (f *Franc) times(multiplier int) Franc {
	return Franc{amount: f.amount * multiplier}
}

func (f *Franc) equals(franc Franc) bool {
	return f.amount == franc.amount
}

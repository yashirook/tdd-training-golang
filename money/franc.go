package money

type Franc struct {
	Money
}

func NewFranc(a int) *Franc {
	return &Franc{
		Money{
			amount: a,
		},
	}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

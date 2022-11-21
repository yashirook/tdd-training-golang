package money

type Franc struct {
	Money
}

func NewFranc(amount int, name string) *Franc {
	return &Franc{
		Money{
			amount: amount,
			name:   name,
		},
	}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount*multiplier, f.name)
}

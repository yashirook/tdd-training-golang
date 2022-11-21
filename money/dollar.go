package money

type Dollar struct {
	Money
}

func NewDollar(amount int, name string) *Dollar {
	return &Dollar{
		Money{
			amount: amount,
			name:   name,
		},
	}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount*multiplier, d.name)
}

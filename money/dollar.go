package money

type Dollar struct {
	Money
}

func NewDollar(a int) *Dollar {
	return &Dollar{
		Money{
			amount: a,
		},
	}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

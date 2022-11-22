package money

type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend, addend Money) Sum {
	return Sum{
		augend: augend,
		addend: addend,
	}
}

func (s Sum) Reduce(currency string) Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, currency)
}

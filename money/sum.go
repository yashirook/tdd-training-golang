package money

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend, addend Expression) Sum {
	return Sum{
		augend: augend,
		addend: addend,
	}
}

func (s Sum) Reduce(bank Bank, currency string) Money {
	amount := s.augend.Reduce(bank, "USD").amount + s.addend.Reduce(bank, "USD").amount
	return NewMoney(amount, currency)
}

func (s Sum) Plus(added Expression) Expression {
	return Sum{}
}

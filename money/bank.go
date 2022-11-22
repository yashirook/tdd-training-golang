package money

type Bank struct{}

func (b Bank) Reduce(source Expression, currency string) Money {
	sum := source.(Sum)
	amount := sum.augend.amount + sum.addend.amount

	return NewMoney(amount, currency)
}

func NewBank() Bank {
	return Bank{}
}

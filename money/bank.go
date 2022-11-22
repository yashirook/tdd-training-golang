package money

type Bank struct{}

func (b Bank) Reduce(source Expression, currency string) Money {
	sum := source.(Sum)
	return sum.Reduce(currency)
}

func NewBank() Bank {
	return Bank{}
}

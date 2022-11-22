package money

type Bank struct{}

func (b Bank) Reduce(source Expression, currency string) Money {
	switch source.(type) {
	case Money:
		return source.(Money)
	}

	sum := source.(Sum)
	return sum.Reduce(currency)
}

func NewBank() Bank {
	return Bank{}
}

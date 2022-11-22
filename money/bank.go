package money

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Reduce(source Expression, currency string) Money {
	return source.Reduce(b, currency)
}

func (b Bank) AddRate(from string, to string, rate int) {
	// if (from == "CHF" && to == "USD" ) {

	// }
}

func (b Bank) Rate(from, to string) int {
	var rate int
	if from == "CHF" && to == "USD" {
		rate = 2
	} else {
		rate = 1
	}
	return rate
}

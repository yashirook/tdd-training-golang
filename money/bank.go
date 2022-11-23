package money

type Bank struct {
	rates map[Pair]int
}

type Pair struct {
	from, to string
}

func NewBank() Bank {
	b := Bank{}
	b.rates = make(map[Pair]int)
	return b
}

func (b *Bank) Reduce(source Expression, currency string) Money {
	return source.Reduce(*b, currency)
}

func (b *Bank) AddRate(from string, to string, rate int) {
	b.rates[Pair{from: from, to: to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}

	return b.rates[Pair{from: from, to: to}]
}

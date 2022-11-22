package money

type Paymenter interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func NewDollar(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "USD",
	}
}

func NewFranc(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "CHF",
	}
}

// 【Check】TDD本では引数をObject型にしていたことについて話す

func (m *Money) Equals(p Paymenter) bool {
	return m.amount == p.Amount() && m.Currency() == p.Currency()
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// 【Check】抽象化を考えるフェーズについて話す

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Plus(p Paymenter) *Money {
	return &Money{
		amount:   m.Amount() + p.Amount(),
		currency: m.Currency(),
	}
}

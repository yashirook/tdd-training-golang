package money

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func NewDollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}

func NewFranc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}

// 【Check】TDD本では引数をObject型にしていたことについて話す

func (m Money) Equals(target Money) bool {
	return m.amount == target.amount && m.currency == target.currency
}

func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// 【Check】抽象化を考えるフェーズについて話す

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Plus(added Money) Expression {
	return NewSum(m, added)
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

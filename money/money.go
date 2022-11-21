package money

type Paymenter interface {
	Amount() int
	Name() string
}

type Money struct {
	amount int
	name   string
}

func NewDollar(amount int) *Money {
	return &Money{
		amount: amount,
		name:   "Dollar",
	}
}

func NewFranc(amount int) *Money {
	return &Money{
		amount: amount,
		name:   "Franc",
	}
}

// 【Check】TDD本では引数をObject型にしていたことについて話す

func (m *Money) Equals(p Paymenter) bool {
	return m.amount == p.Amount() && m.Name() == p.Name()
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{
		amount: m.amount * multiplier,
		name:   m.name,
	}
}

// 【Check】抽象化を考えるフェーズについて話す

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Name() string {
	return m.name
}

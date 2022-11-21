package money

type Paymenter interface {
	Amount() int
	Name() string
}

type Money struct {
	amount int
	name   string
}

// 【Check】TDD本では引数をObject型にしていたことについて話す

func (m *Money) Equals(p Paymenter) bool {
	return m.amount == p.Amount() && m.Name() == p.Name()
}

// 【Check】抽象化を考えるフェーズについて話す

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Name() string {
	return m.name
}

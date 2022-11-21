package money

type Paymenter interface {
	Amount() int
}

type Money struct {
	amount int
}

// 【Check】TDD本では引数をObject型にしていたことについて話す

func (m *Money) Equals(p Paymenter) bool {
	return m.amount == p.Amount()
}

// 【Check】抽象化を考えるフェーズについて話す

func (m *Money) Amount() int {
	return m.amount
}

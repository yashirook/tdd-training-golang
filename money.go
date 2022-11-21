package main

type Money struct {
	amount int
}

type Doller = Money
type Franc = Money

func (m *Money) times(multipulier int) Money {
	return Money{m.amount * multipulier}
}

// 【Check】TDD本では引数をObject型にしていたことについて話す
func (m *Money) equals(money Money) bool {
	return m.amount == money.amount
}

// 【Check】抽象化を考えるフェーズについて話す

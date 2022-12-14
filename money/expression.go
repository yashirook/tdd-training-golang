package money

type Expression interface {
	Reduce(bank Bank, to string) Money
	Plus(added Expression) Expression
	Times(multiplier int) Expression
}

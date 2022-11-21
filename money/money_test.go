package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	assert.Equal(t, &Money{amount: 10, currency: "USD"}, five.Times(2))
	assert.Equal(t, &Money{amount: 15, currency: "USD"}, five.Times(3))
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, five.Equals(NewDollar(5)))
	assert.False(t, five.Equals(NewDollar(6)))
	assert.False(t, five.Equals(NewFranc(5)))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, &Money{amount: 10, currency: "CHF"}, five.Times(2))
	assert.Equal(t, &Money{amount: 15, currency: "CHF"}, five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	five := NewFranc(5)
	assert.True(t, five.Equals(NewFranc(5)))
	assert.False(t, five.Equals(NewFranc(6)))
	assert.False(t, five.Equals(NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}
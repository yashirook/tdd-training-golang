package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	assert.Equal(t, &Dollar{Money{amount: 10}}, five.Times(2))
	assert.Equal(t, &Dollar{Money{amount: 15}}, five.Times(3))
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, five.Equals(NewDollar(5)))
	assert.False(t, five.Equals(NewDollar(6)))
	// assert.False(t, five.Equals(NewDollar(5)))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, &Franc{Money{amount: 10}}, five.Times(2))
	assert.Equal(t, &Franc{Money{amount: 15}}, five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	five := NewFranc(5)
	assert.True(t, five.Equals(NewFranc(5)))
	assert.False(t, five.Equals(NewFranc(6)))
}

package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5, "Dollar")

	assert.Equal(t, &Dollar{Money{amount: 10, name: "Dollar"}}, five.Times(2))
	assert.Equal(t, &Dollar{Money{amount: 15, name: "Dollar"}}, five.Times(3))
}

func TestEquality(t *testing.T) {
	five := NewDollar(5, "Dollar")
	assert.True(t, five.Equals(NewDollar(5, "Dollar")))
	assert.False(t, five.Equals(NewDollar(6, "Dollar")))
	assert.False(t, five.Equals(NewDollar(5, "Franc")))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5, "Franc")
	assert.Equal(t, &Franc{Money{amount: 10, name: "Franc"}}, five.Times(2))
	assert.Equal(t, &Franc{Money{amount: 15, name: "Franc"}}, five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	five := NewFranc(5, "Franc")
	assert.True(t, five.Equals(NewFranc(5, "Franc")))
	assert.False(t, five.Equals(NewFranc(6, "Franc")))
	assert.False(t, five.Equals(NewFranc(5, "Dollar")))
}

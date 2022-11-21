package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := &Doller{5}

	product := five.times(2)
	assert.Equal(t, 10, product.amount)

	product = five.times(3)
	assert.Equal(t, 15, product.amount)
}

func TestEquality(t *testing.T) {
	five := &Doller{5}
	assert.True(t, five.equals(&Doller{5}))
	assert.False(t, five.equals(&Doller{6}))
}

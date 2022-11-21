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
	assert.True(t, &Doller{5}.equals(&Doller{5}))
}

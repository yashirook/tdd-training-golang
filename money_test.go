package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Doller{5}

	product := five.times(2)
	assert.Equal(t, Doller{10}, product)

	product = five.times(3)
	assert.Equal(t, Doller{15}, product)
}

func TestEquality(t *testing.T) {
	five := Doller{5}
	assert.True(t, five.equals(Doller{5}))
	assert.False(t, five.equals(Doller{6}))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := &Doller{5}

	five.times(2)
	assert.Equal(t, 10, five.amount)
}

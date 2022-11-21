package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Doller{5}

	assert.Equal(t, Doller{10}, five.times(2))
	assert.Equal(t, Doller{15}, five.times(3))
}

func TestEquality(t *testing.T) {
	five := Doller{5}
	assert.True(t, five.equals(Doller{5}))
	assert.False(t, five.equals(Doller{6}))
}

func TestFrancMultiplication(t *testing.T) {
	five := Franc{5}
	assert.Equal(t, Franc{10}, five.times(2))
	assert.Equal(t, Franc{15}, five.times(3))
}

func TestFrancEquality(t *testing.T) {
	five := Franc{5}
	assert.True(t, five.equals(Franc{5}))
	assert.False(t, five.equals(Franc{6}))
}

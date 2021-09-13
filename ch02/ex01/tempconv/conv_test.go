package tempconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCToF(t *testing.T) {
	input := Celsius(100)
	actual := CToF(input)
	assert.Equal(t, Fahrenheit(212.0), actual)
}

func TestCToK(t *testing.T) {
	input := Celsius(100)
	actual := CToK(input)
	assert.Equal(t, Kelvin(373.15), actual)
}

func TestFToC(t *testing.T) {
	input := Fahrenheit(212.0)
	actual := FToC(input)
	assert.Equal(t, Celsius(100.0), actual)
}

func TestFToK(t *testing.T) {
	input := Fahrenheit(212.0)
	actual := FToK(input)
	assert.Equal(t, Kelvin(373.15), actual)
}

func TestKToC(t *testing.T) {
	input := Kelvin(274.15)
	actual := KToC(input)
	assert.Equal(t, Celsius(1), actual)
}

func TestKToF(t *testing.T) {
	input := Kelvin(273.15)
	actual := KToF(input)
	assert.Equal(t, Fahrenheit(32), actual)
}

package tempconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCToF(t *testing.T) {
	temperatureC := Celsius(100)
	actualF := CToF(temperatureC)
	assert.Equal(t, Fahrenheit(212.0), actualF)
}

func TestFToC(t *testing.T) {
	temperatureF := Fahrenheit(212.0)
	actualC := FToC(temperatureF)
	assert.Equal(t, Celsius(100.0), actualC)
}

package tempconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCelsius_String(t *testing.T) {
	temperature := Celsius(100)
	actual := temperature.String()
	assert.Equal(t, "100°C", actual)
}

func TestFahrenheit_String(t *testing.T) {
	temperature := Fahrenheit(100)
	actual := temperature.String()
	assert.Equal(t, "100°F", actual)
}

func TestKelvin_String(t *testing.T) {
	temperature := Kelvin(100)
	actual := temperature.String()
	assert.Equal(t, "100K", actual)
}

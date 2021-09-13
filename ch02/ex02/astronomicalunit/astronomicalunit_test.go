package astronomicalunit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAstronomicalUnit_String(t *testing.T) {
	input := AstronomicalUnit(100)
	actual := input.String()
	assert.Equal(t, "100 au", actual)
}

func TestLightYear_String(t *testing.T) {
	input := LightYear(100)
	actual := input.String()
	assert.Equal(t, "100 ly", actual)
}

func TestParsec_String(t *testing.T) {
	input := Parsec(100)
	actual := input.String()
	assert.Equal(t, "100 pc", actual)
}

func TestMeter_String(t *testing.T) {
	input := Meter(100)
	actual := input.String()
	assert.Equal(t, "100 m", actual)
}

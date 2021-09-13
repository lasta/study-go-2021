package astronomicalunit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeter_ToAU(t *testing.T) {
	input := Meter(1.0)
	actual := input.ToAU()
	assert.InDelta(t, 6.68459e-12, float64(actual), 1.0e-17)
}

func TestMeter_ToLY(t *testing.T) {
	input := Meter(1.0)
	actual := input.ToLY()
	assert.InDelta(t, 1.05700e-16, float64(actual), 1.0e-22)
}

func TestMeter_ToPC(t *testing.T) {
	input := Meter(1.0)
	actual := input.ToPC()
	assert.InDelta(t, 3.24078e-17, float64(actual), 1.0e-22)
}

func TestAstronomicalUnit_ToMeter(t *testing.T) {
	input := AstronomicalUnit(2.0)
	actual := input.ToMeter()
	assert.InDelta(t, 2.99196e11, float64(actual), 1.0e5)
}

func TestLightYear_ToMeter(t *testing.T) {
	input := LightYear(2.0)
	actual := input.ToMeter()
	assert.InDelta(t, 1.892146e16, float64(actual), 1.0e10)
}

func TestParsec_ToMeter(t *testing.T) {
	input := Parsec(2.0)
	actual := input.ToMeter()
	assert.InDelta(t, 6.17136e16, float64(actual), 1.0e10)
}

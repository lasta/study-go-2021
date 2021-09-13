package astronomicalunit

import "fmt"

type Meter float64
type AstronomicalUnit float64
type LightYear float64
type Parsec float64

const (
	AUInMeter Meter = 1.49598e11
	LYInMeter Meter = 9.46073e15
	PCInMeter Meter = 3.08568e16
)

func (au AstronomicalUnit) String() string {
	return fmt.Sprintf("%g au", au)
}

func (ly LightYear) String() string {
	return fmt.Sprintf("%g ly", ly)
}

func (pc Parsec) String() string {
	return fmt.Sprintf("%g pc", pc)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

package tempconv0

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroCelsius Celsius = -273.15
	FreezingCelsius     Celsius = 0.0
	BoilingCelsius      Celsius = 100.0
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

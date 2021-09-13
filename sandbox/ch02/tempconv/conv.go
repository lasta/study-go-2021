package tempconv

// CToF converts from degree Celsius to degree Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts from degree Fahrenheit to degree Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

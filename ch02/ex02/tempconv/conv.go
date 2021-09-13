package tempconv

// CToF converts from degree Celsius to degree Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts from degree Celsius to degree Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroCelsius)
}

// FToC converts from degree Fahrenheit to degree Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK converts from degree Fahrenheit to degree Kelvin
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

// KToC converts from degree Kelvin to degree Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroCelsius
}

// KToF converts from degree Kelvin to degree Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

package astronomicalunit

func (m Meter) ToAU() AstronomicalUnit {
	return AstronomicalUnit(m / AUInMeter)
}

func (m Meter) ToLY() LightYear {
	return LightYear(m / LYInMeter)
}

func (m Meter) ToPC() Parsec {
	return Parsec(m / PCInMeter)
}

func (au AstronomicalUnit) ToMeter() Meter {
	return Meter(au) * AUInMeter
}

func (ly LightYear) ToMeter() Meter {
	return Meter(ly) * LYInMeter
}

func (pc Parsec) ToMeter() Meter {
	return Meter(pc) * PCInMeter
}

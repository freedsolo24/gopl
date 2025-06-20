package tempconv

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func MtoMI(m Meter) Mile {
	return Mile(m / 1609.344)
}

func MItoM(mi Mile) Meter {
	return Meter(mi * 1609.344)
}

func PoundtoKilogram(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

func KilogramtoPound(k Kilogram) Pound {
	return Pound(k * 2.20462262)
}

func angleClock(hour int, minutes int) float64 {
	a := float64(hour%12)*30.0 + float64(minutes)*0.5
	b := float64(minutes) * 6.0
	dif := math.Abs(a - b)
	return math.Min(dif, 360.0-dif)
}
func internalAngles(sides []int) []float64 {
	s := []float64{float64(sides[0]), float64(sides[1]), float64(sides[2])}
	sort.Float64s(s)
	if s[0]+s[1] <= s[2] {
		return []float64{}
	}
	a1, a2, a3 := h1(s[0], s[1], s[2]), h1(s[1], s[0], s[2]), h1(s[2], s[0], s[1])
	return []float64{a1, a2, a3}
}

func h1(opp, x, y float64) float64 {
	a := (x*x + y*y - opp*opp) / (2.0 * x * y)
	if a < -1.0 {
		a = -1.0
	} else if a > 1.0 {
		a = 1.0
	}
	return math.Acos(a) * 180.0 / math.Pi
}
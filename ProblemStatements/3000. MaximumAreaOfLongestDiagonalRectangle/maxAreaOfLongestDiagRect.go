func areaOfMaxDiagonal(dimensions [][]int) int {
	a, b := dimensions[0][0]*dimensions[0][0]+dimensions[0][1]*dimensions[0][1], dimensions[0][0]*dimensions[0][1]
	for i := 1; i < len(dimensions); i++ {
		temp := dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]
		if temp < a {
			continue
		}
		area := dimensions[i][0] * dimensions[i][1]
		if temp > a {
			a = temp
			b = area
		} else {
			b = max(b, area)
		}
	}
	return b
}
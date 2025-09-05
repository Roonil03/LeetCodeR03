func findClosest(x int, y int, z int) int {
	a, b := abs(z-x), abs(z-y)
	if a > b {
		return 2
	} else if b > a {
		return 1
	} else {
		return 0
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
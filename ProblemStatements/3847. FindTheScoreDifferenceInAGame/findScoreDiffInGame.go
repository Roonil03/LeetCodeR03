func scoreDifference(nums []int) int {
	fg, res := true, 0
	for i, v := range nums {
		if v%2 != 0 {
			fg = !fg
		}
		if (1+i)%6 == 0 {
			fg = !fg
		}
		if fg {
			res += v
		} else {
			res -= v
		}
	}
	return res
}
func threeConsecutiveOdds(arr []int) bool {
	res := 0
	for _, i := range arr {
		if i%2 != 0 {
			res++
			if res == 3 {
				return true
			}
		} else {
			res = 0
		}
	}
	return false
}
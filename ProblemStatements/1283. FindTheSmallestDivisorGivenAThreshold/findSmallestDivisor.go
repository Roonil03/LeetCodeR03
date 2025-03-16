func smallestDivisor(nums []int, threshold int) int {
	helper := func(div int) bool {
		sum := 0
		for _, num := range nums {
			sum += (num-1)/div + 1
		}
		return sum <= threshold
	}
	a, b := 1, 0
	for _, num := range nums {
		if num > b {
			b = num
		}
	}
	for a < b {
		m := a + (b-a)/2
		if helper(m) {
			b = m
		} else {
			a = m + 1
		}
	}
	return a
}
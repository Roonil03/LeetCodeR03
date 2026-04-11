func countDigitOccurrences(nums []int, digit int) int {
	res := 0
	for _, v := range nums {
		x := v
		for x > 0 {
			if x%10 == digit {
				res++
			}
			x /= 10
		}
	}
	return res
}
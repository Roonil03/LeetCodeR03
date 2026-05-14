func isGood(nums []int) bool {
	n := len(nums) - 1
	if n < 1 {
		return false
	}
	c := make([]int, n+1)
	for _, v := range nums {
		if v < 1 || v > n {
			return false
		}
		c[v]++
	}
	for i := 1; i < n; i++ {
		if c[i] != 1 {
			return false
		}
	}
	return c[n] == 2
}
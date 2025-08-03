func longestSubarray(nums []int) int {
	a := nums[0]
	for _, n := range nums {
		if n > a {
			a = n
		}
	}
	cur, res := 0, 0
	for _, n := range nums {
		if n == a {
			cur++
			if cur > res {
				res = cur
			}
		} else {
			cur = 0
		}
	}
	return res
}
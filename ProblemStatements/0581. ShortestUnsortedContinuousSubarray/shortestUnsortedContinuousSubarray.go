func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	mm, nm, l, r := nums[0], nums[n-1], -1, -2
	for i := range n {
		if nums[i] < mm {
			r = i
		} else {
			mm = nums[i]
		}
		if j := n - 1 - i; nums[j] > nm {
			l = j
		} else {
			nm = nums[j]
		}
	}
	return r - l + 1
}
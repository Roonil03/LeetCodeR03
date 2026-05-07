func maxValue(nums []int) []int {
	n := len(nums)
	pre, suf, res := make([]int, n), make([]int, n), make([]int, n)
	pre[0] = nums[0]
	suf[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		pre[i] = max(nums[i], pre[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(nums[i], suf[i+1])
	}
	res[n-1] = pre[n-1]
	for i := n - 2; i >= 0; i-- {
		res[i] = pre[i]
		if pre[i] > suf[i+1] {
			res[i] = res[i+1]
		}
	}
	return res
}
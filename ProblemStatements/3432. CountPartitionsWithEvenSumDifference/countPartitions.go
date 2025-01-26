func countPartitions(nums []int) int {
	p := make([]int, len(nums))
	n := len(nums)
	p[0] = nums[0]
	for i := 1; i < n; i++ {
		p[i] = p[i-1] + nums[i]
	}
	res := 0
	for i := 0; i < n-1; i++ {
		if ((p[n-1]-p[i])-p[i])%2 == 0 {
			res++
		}
	}
	return res
}
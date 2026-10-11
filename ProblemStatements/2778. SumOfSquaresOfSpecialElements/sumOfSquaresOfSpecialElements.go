func sumOfSquares(nums []int) int {
	n := len(nums)
	res := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res += nums[i-1] * nums[i-1]
			if i*i != n {
				res += nums[n/i-1] * nums[n/i-1]
			}
		}
	}
	return res
}
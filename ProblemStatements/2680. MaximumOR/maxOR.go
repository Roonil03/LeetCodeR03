func maximumOr(nums []int, k int) int64 {
	// n := len(nums)
	// if n == 0 {
	// 	return 0
	// }
	// prefix := make([]int, n+1)
	// suffix := make([]int, n+1)
	// for i := 1; i < n; i++ {
	// 	prefix[i] = prefix[i-1] | nums[i-1]
	// 	suffix[n-i-1] = suffix[n-i] | nums[n-i]
	// }
	// var res int64 = 0
	// for i := 0; i < n; i++ {
	// 	c := int64(prefix[i]) | (int64(nums[i]) << k) | int64(suffix[i])
	// 	if c > res {
	// 		res = c
	// 	}
	// }
	// return res

	var sing, mult int
	for _, x := range nums {
		mult |= x & sing
		sing |= x
	}
	best := sing
	for _, x := range nums {
		rem := mult | (sing &^ x)
		add := rem | (x << k)
		best = max(best, add)
	}
	return int64(best)
}
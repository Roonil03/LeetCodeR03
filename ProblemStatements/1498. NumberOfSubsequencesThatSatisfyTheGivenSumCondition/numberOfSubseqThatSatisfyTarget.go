func numSubseq(nums []int, target int) int {
	const mod = 1e9 + 7
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}
	sort.Ints(nums)
	res := 0
	l, r := 0, n-1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			res = (res + pow2[r-l]) % mod
			l++
		}
	}
	return res
}
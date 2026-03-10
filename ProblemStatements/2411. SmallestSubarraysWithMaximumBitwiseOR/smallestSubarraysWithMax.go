func smallestSubarrays(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	b := 32
	pos := make([]int, b)
	for i := 0; i < b; i++ {
		pos[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		a := nums[i]
		for j := 0; j < b; j++ {
			if a&(1<<j) != 0 {
				pos[j] = i
			}
		}
		k := i
		for _, p := range pos {
			if p < n && p > k {
				k = p
			}
		}
		res[i] = k - i + 1
	}
	return res
}
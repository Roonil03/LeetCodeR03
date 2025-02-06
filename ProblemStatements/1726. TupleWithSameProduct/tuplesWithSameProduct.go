func tupleSameProduct(nums []int) int {
	n := len(nums)
	m := make(map[int]int)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			prod := nums[i] * nums[j]
			res += 8 * m[prod]
			m[prod]++
		}
	}
	return res
}
func countGood(nums []int, k int) int64 {
	var res int64
	res = 0
	m := make(map[int]int)
	for i, j := 0, 0; j < len(nums); j++ {
		k -= m[nums[j]]
		m[nums[j]]++
		for k <= 0 {
			m[nums[i]]--
			k += m[nums[i]]
			i++
		}
		res += int64(i)
	}
	return res
}
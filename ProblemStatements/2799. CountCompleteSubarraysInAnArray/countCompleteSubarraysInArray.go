func countCompleteSubarrays(nums []int) int {
	n, s, res, j, m := len(nums), 0, 0, 0, make(map[int]int)
	for _, num := range nums {
		if m[num] == 0 {
			s++
		}
		m[num]++
	}
	m = make(map[int]int)
	i := 0
	for j = 0; j < n; j++ {
		if m[nums[j]] == 0 {
			s--
		}
		m[nums[j]]++
		for s == 0 {
			m[nums[i]]--
			if m[nums[i]] == 0 {
				s++
			}
			i++
		}
		res += i
	}
	return res
}
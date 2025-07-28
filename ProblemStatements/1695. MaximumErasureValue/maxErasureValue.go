func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]bool)
	l, r := 0, 0
	temp, res := 0, 0
	for r < len(nums) {
		if !m[nums[r]] {
			m[nums[r]] = true
			temp += nums[r]
			res = max(res, temp)
			r++
		} else {
			m[nums[l]] = false
			temp -= nums[l]
			l++
		}
	}
	return res
}
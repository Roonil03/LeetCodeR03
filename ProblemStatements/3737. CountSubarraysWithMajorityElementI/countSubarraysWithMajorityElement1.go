func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	count := make([]int, 2*n+2)
	o := n + 1
	count[o] = 1
	s, sm, res := 0, 0, 0
	for _, v := range nums {
		if v == target {
			sm += count[s+o]
			s++
		} else {
			s--
			sm -= count[s+o]
		}
		res += sm
		count[s+o]++
	}
	return res
}
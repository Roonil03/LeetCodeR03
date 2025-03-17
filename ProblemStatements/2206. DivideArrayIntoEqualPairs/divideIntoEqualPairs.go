func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	for _, c := range m {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
func maxOperations(nums []int, k int) int {
	count := 0
	m := make(map[int]int)
	for _, num := range nums {
		a := k - num
		if m[a] > 0 {
			count++
			m[a]--
		} else {
			m[num]++
		}
	}
	return count
}
func pivotIndex(nums []int) int {
	tot := 0
	for _, v := range nums {
		tot += v
	}
	l := 0
	for i, v := range nums {
		if l == tot-l-v {
			return i
		}
		l += v
	}
	return -1
}
func containsNearbyDuplicate(nums []int, k int) bool {
	a := make(map[int]int)
	for i, el := range nums {
		l, exists := a[el]
		if exists && i-l <= k {
			return true
		}
		a[el] = i
	}
	return false
}

func increasingTriplet(nums []int) bool {
	a, b := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, n := range nums {
		if n <= a {
			a = n
		} else if n <= b {
			b = n
		} else {
			return true
		}
	}
	return false
}
func maximumCount(nums []int) int {
	a := binarySearch(nums, 0)
	b := len(nums) - binarySearch(nums, 1)
	return int(math.Max(float64(a), float64(b)))
}

func binarySearch(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	res := len(nums)
	for l <= r {
		m := (l + r) / 2
		if nums[m] < k {
			l = m + 1
		} else {
			res = m
			r = m - 1
		}
	}
	return res
}
func countSubarrays(nums []int, k int) int64 {
	var n, count, l, r, res int64 = maxArray(nums), 0, 0, 0, 0
	for r < int64(len(nums)) {
		if int64(nums[r]) == n {
			count++
		}
		for count >= int64(k) {
			if int64(nums[l]) == n {
				count--
			}
			l++
			res += int64(len(nums)) - r
		}
		r++
	}
	return res
}

func maxArray(nums []int) int64 {
	res := math.MinInt32
	for _, i := range nums {
		if res < i {
			res = i
		}
	}
	return int64(res)
}
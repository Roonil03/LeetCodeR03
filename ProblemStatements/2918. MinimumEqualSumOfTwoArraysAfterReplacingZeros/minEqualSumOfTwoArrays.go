func minSum(nums1 []int, nums2 []int) int64 {
	zero1, zero2, sum1, sum2 := int64(0), int64(0), int64(0), int64(0)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == 0 {
			zero1++
		}
		sum1 += int64(nums1[i])
	}
	for i := 0; i < len(nums2); i++ {
		if nums2[i] == 0 {
			zero2++
		}
		sum2 += int64(nums2[i])
	}
	if zero1 > 0 && zero2 > 0 {
		if sum1+zero1 >= sum2+zero2 {
			return sum1 + zero1
		}
		return sum2 + zero2
	}
	if zero1 > 0 {
		if sum1+zero1 <= sum2 {
			return sum2
		}
		return -1
	}
	if zero2 > 0 {
		if sum2+zero2 <= sum1 {
			return sum1
		}
		return -1
	}
	if sum1 == sum2 {
		return sum1
	}
	return -1
}
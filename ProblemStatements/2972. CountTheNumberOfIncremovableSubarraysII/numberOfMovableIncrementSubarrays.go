func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	leftIdx := 0

	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			leftIdx++
		} else {
			break
		}
	}

	if leftIdx == n-1 {
		ans := int64(n * (n + 1) / 2)
		return ans
	}

	rightIdx := int64(n - 1)

	for i := n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			rightIdx--
		} else {
			break
		}
	}

	totalIncremovableSubarrays := int64(0)
	totalIncremovableSubarrays += int64(n - int(rightIdx) + 1)

	l, r := 0, int(rightIdx)

	for l <= leftIdx {
		for r < n && nums[l] >= nums[r] {
			r++
		}
		totalIncremovableSubarrays += int64(n - r + 1)
		l++
	}

	return totalIncremovableSubarrays
}
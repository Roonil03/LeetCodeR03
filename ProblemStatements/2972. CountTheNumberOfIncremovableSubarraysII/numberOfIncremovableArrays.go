func incremovableSubarrayCount(nums []int) int64 {
	nums = append([]int{0}, append(nums, int(^uint(0)>>1))...)
	n := len(nums)
	i := 0
	for i < n-1 {
		if nums[i] >= nums[i+1] {
			break
		}
		i++
	}
	if i == n-1 {
		return int64((n - 2) * (n - 1) / 2)
	}
	j := n - 1
	for j > 0 {
		if nums[j-1] >= nums[j] {
			break
		}
		j--
	}
	l := 0
	r := j
	count := int64(0)
	for l <= i {
		for r < n && nums[l] >= nums[r] {
			r++
		}
		count += int64(n - r)
		l++
	}
	return count
}
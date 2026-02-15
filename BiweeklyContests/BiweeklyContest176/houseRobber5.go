func rob(nums []int, colors []int) int64 {
	t1, t2 := int64(nums[0]), int64(0)
	for i := 1; i < len(nums); i++ {
		temp := int64(0)
		if colors[i] == colors[i-1] {
			temp = t2 + int64(nums[i])
		} else {
			temp = max(t1, t2) + int64(nums[i])
		}
		g := max(t1, t2)
		t1 = temp
		t2 = g
	}
	return max(t1, t2)
}
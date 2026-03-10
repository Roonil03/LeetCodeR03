func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		id := abs(nums[i]) - 1
		if nums[id] > 0 {
			nums[id] = -nums[id]
		}
	}
	var res []int
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
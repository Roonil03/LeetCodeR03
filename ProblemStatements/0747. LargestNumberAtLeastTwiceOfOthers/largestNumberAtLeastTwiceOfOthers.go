func dominantIndex(nums []int) int {
	res, id, b := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > res {
			b = res
			res = nums[i]
			id = i
		} else if nums[i] > b {
			b = nums[i]
		}
	}
	if b*2 <= res {
		return id
	}
	return -1
}
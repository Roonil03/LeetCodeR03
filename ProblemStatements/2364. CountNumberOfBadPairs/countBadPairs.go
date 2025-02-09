func countBadPairs(nums []int) int64 {
	temp := make(map[int]int)
	var res int64
	res = 0
	for i := 0; i < len(nums); i++ {
		dif := nums[i] - i
		if c, exists := temp[dif]; exists {
			res += int64(i - c)
			temp[dif]++
		} else {
			res += int64(i)
			temp[dif] = 1
		}
	}
	return res
}
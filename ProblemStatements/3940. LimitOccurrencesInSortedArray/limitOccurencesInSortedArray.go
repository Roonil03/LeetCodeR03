func limitOccurrences(nums []int, k int) []int {
	w := k
	for _, v := range nums[k:] {
		if v != nums[w-k] {
			nums[w] = v
			w++
		}
	}
	return nums[:w]
}
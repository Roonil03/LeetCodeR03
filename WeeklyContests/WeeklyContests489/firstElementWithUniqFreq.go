func firstUniqueFreq(nums []int) int {
	mp := make(map[int]int)
	for _, i := range nums {
		mp[i]++
	}
	freq := make(map[int]int)
	for _, i := range mp {
		freq[i]++
	}
	for _, i := range nums {
		f := mp[i]
		if freq[f] == 1 {
			return i
		}
	}
	return -1
}
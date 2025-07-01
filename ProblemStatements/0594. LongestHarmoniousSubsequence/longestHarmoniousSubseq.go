func findLHS(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res := 0
	for num, c := range freq {
		if v, ok := freq[num+1]; ok {
			if c+v > res {
				res = c + v
			}
		}
	}
	return res
}
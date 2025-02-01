func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := 0
	for _, num := range nums {
		if n < num {
			n = num
		}
		res += n - num
		n++
	}
	return res
}
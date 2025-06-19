func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	res := 1
	s := nums[0]
	for _, n := range nums {
		if n-s > k {
			res++
			s = n
		}
	}
	return res
}
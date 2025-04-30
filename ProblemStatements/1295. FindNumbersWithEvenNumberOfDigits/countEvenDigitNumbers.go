func findNumbers(nums []int) int {
	res := 0
	for _, i := range nums {
		if (i >= 10 && i <= 99) || (i > 999 && i < 10000) || i == 100000 {
			res++
		}
	}
	return res
}
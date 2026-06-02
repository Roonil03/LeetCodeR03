func minimumSwaps(nums []int) int {
	a := 0
	for _, v := range nums {
		if v == 0 {
			a++
		}
	}
	res := 0
	for _, v := range nums[:len(nums)-a] {
		if v == 0 {
			res++
		}
	}
	return res
}
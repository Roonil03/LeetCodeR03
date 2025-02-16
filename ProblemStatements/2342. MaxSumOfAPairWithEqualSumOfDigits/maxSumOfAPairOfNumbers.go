func maximumSum(nums []int) int {
	res := -1
	a := make(map[int]int)
	for _, i := range nums {
		d := sum(i)
		if a[d] != 0 {
			res = max(res, i+a[d])
		}
		a[d] = max(a[d], i)
	}
	return res
}

func sum(a int) int {
	res := 0
	for a > 0 {
		res += a % 10
		a /= 10
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
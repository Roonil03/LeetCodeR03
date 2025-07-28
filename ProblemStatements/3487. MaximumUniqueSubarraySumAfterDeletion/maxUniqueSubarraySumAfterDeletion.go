func maxSum(nums []int) int {
	u := map[int]int{}
	temp := -math.MaxInt
	res, a := 0, false
	for _, n := range nums {
		if _, ok := u[n]; !ok {
			if n >= 0 {
				u[n]++
				res += n
				a = true
			} else {
				temp = max(temp, n)
			}
		}
	}
	if !a {
		return temp
	}
	return res
}
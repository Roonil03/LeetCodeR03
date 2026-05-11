func separateDigits(nums []int) []int {
	var res []int
	for _, v := range nums {
		s := strconv.Itoa(v)
		for _, c := range s {
			res = append(res, int(c-'0'))
		}
	}
	return res
}
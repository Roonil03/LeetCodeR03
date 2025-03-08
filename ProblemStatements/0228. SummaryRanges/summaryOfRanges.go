func summaryRanges(nums []int) []string {
	var res []string
	n := len(nums)
	if n == 0 {
		return res
	}

	for i := 0; i < n; {
		a := nums[i]
		b := a
		for i < n-1 && nums[i+1] == b+1 {
			b = nums[i+1]
			i++
		}
		if a == b {
			res = append(res, strconv.Itoa(a))
		} else {
			res = append(res, strconv.Itoa(a)+"->"+strconv.Itoa(b))
		}
		i++
	}
	return res
}
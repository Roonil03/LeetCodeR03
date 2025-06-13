func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]-nums[0]
	for l < r {
		m := l + (r-l)/2
		if h1(nums, m) >= p {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func h1(nums []int, d int) int {
	p, i := 0, 0
	for i < len(nums)-1 {
		if nums[i+1]-nums[i] <= d {
			p++
			i += 2
		} else {
			i++
		}
	}
	return p
}
func maxDigitRange(nums []int) int {
	temp, res := -1, 0
	for _, v := range nums {
		n := v
		mi := n % 10
		ma := mi
		for n /= 10; n > 0; n /= 10 {
			d := n % 10
			if d < mi {
				mi = d
			}
			if d > ma {
				ma = d
			}
			if ma-mi == 9 {
				break
			}
		}
		r := ma - mi
		if r > temp {
			temp = r
			res = v
		} else if r == temp {
			res += v
		}
	}
	return res
}
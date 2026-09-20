func isPossible(nums []int) bool {
	p1, p2, p3, prev := 0, 0, 0, 0
	for i := 0; i < len(nums); {
		cur, c := nums[i], 0
		for i < len(nums) && nums[i] == cur {
			c++
			i++
		}
		if cur != prev+1 {
			if p1 > 0 || p2 > 0 {
				return false
			}
			p1, p2, p3 = c, 0, 0
		} else {
			if c < p1+p2 {
				return false
			}
			c1, c2, c3 := c-p1-p2, p1, p2
			if c1 > p3 {
				c3 += p3
				c1 -= p3
			} else {
				c3 += c1
				c1 = 0
			}
			p1, p2, p3 = c1, c2, c3
		}
		prev = cur
	}
	return p1 == 0 && p2 == 0
}
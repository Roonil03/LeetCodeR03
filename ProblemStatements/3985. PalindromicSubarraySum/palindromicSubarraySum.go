func getSum(nums []int) int64 {
	n := len(nums)
	p := make([]int, n+1)
	for i := range n {
		p[i+1] = p[i] + nums[i]
	}
	a := bits.Len(uint(n + 1))
	st := make([][]int, n+1)
	for i := range n + 1 {
		st[i] = make([]int, a)
		st[i][0] = p[i]
	}
	for i := 1; i < a; i++ {
		for j := 0; j+(1<<i) <= n+1; j++ {
			x := st[j][i-1]
			y := st[j+(1<<(i-1))][i-1]
			st[j][i] = max(x, y)
		}
	}
	query := func(l, r int) int {
		j := bits.Len(uint(r-l+1)) - 1
		x := st[l][j]
		y := st[r-(1<<j)+1][j]
		return max(x, y)
	}
	rad := make([]int, 2*n+1)
	c, r := 0, 0
	for i := range 2*n + 1 {
		ri := 0
		if i < r {
			m := 2*c - i
			ri = rad[m]
			if i+ri > r {
				ri = r - i
			}
		}
		for i-ri-1 >= 0 && i+ri+1 < 2*n+1 {
			l, h := i-ri-1, i+ri+1
			if l%2 == 0 {
				ri++
			} else if nums[l/2] == nums[h/2] {
				ri++
			} else {
				break
			}
		}
		rad[i] = ri
		if i+ri > r {
			c, r = i, i+ri
		}
	}
	var res int64 = math.MinInt64
	for i := range 2*n + 1 {
		if i%2 == 1 {
			pos := i / 2
			l := pos + 1
			h := pos + (rad[i]-1)/2 + 1
			if l <= h {
				cand := int64(2*query(l, h) - p[pos] - p[pos+1])
				res = max(res, cand)
			}
		} else {
			pos := i / 2
			l := pos + 1
			h := pos + rad[i]/2
			if l <= h {
				cand := int64(2*query(l, h) - 2*p[pos])
				res = max(res, cand)
			}
		}
	}
	return res
}
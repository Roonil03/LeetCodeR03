type hash struct {
	h1 uint64
	h2 uint64
}

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return n
	}
	b1 := uint64(1e9 + 7)
	b2 := uint64(1e9 + 9)
	p1 := make([]uint64, n+1)
	p2 := make([]uint64, n+1)
	p1[0], p2[0] = 1, 1
	for i := 1; i <= n; i++ {
		p1[i] = p1[i-1] * b1
		p2[i] = p2[i-1] * b2
	}
	a1 := func(l1 int) bool {
		h := make([]hash, n-l1+1)
		var cur1, cur2 uint64
		for i := 0; i < l1; i++ {
			cur1 = cur1*b1 + uint64(nums[i])
			cur2 = cur2*b2 + uint64(nums[i])
		}
		h[0] = hash{cur1, cur2}
		pl1 := p1[l1]
		pl2 := p2[l1]
		for i := l1; i < n; i++ {
			cur1 = cur1*b1 - uint64(nums[i-l1])*pl1 + uint64(nums[i])
			cur2 = cur2*b2 - uint64(nums[i-l1])*pl2 + uint64(nums[i])
			h[i-l1+1] = hash{cur1, cur2}
		}
		slices.SortFunc(h, func(a, b hash) int {
			if c := cmp.Compare(a.h1, b.h1); c != 0 {
				return c
			}
			return cmp.Compare(a.h2, b.h2)
		})
		for i := 0; i < len(h); i++ {
			fg := true
			if i > 0 && h[i] == h[i-1] {
				fg = false
			}
			if i < len(h)-1 && h[i] == h[i+1] {
				fg = false
			}
			if fg {
				return true
			}
		}
		return false
	}
	l, r := 1, n
	res := n
	for l <= r {
		mid := l + (r-l)/2
		if a1(mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
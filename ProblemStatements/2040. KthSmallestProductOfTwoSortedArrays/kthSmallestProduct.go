func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	a, b := nums1, nums2
	m := len(b)
	lo, hi := int64(-1e10-1), int64(1e10+1)
	f := func(x int64) int64 {
		var r int64
		for _, v := range a {
			if v == 0 {
				if x >= 0 {
					r += int64(m)
				}
			} else if v > 0 {
				l, h := 0, m
				for l < h {
					mi := (l + h) >> 1
					if int64(v)*int64(b[mi]) <= x {
						l = mi + 1
					} else {
						h = mi
					}
				}
				r += int64(l)
			} else {
				l, h := 0, m
				for l < h {
					mi := (l + h) >> 1
					if int64(v)*int64(b[mi]) <= x {
						h = mi
					} else {
						l = mi + 1
					}
				}
				r += int64(m - l)
			}
		}
		return r
	}
	for lo < hi {
		mi := lo + (hi-lo)/2
		if f(mi) < k {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}
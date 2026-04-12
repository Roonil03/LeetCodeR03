func longestBalanced(s string) int {
	n := len(s)
	t0 := 0
	for i := range s {
		if s[i] == '0' {
			t0++
		}
	}
	t1 := n - t0
	off := n + 2
	sz := 2*n + 5
	fa, fz, fo := make([]int, sz), make([]int, sz), make([]int, sz)
	for i := range sz {
		fa[i], fz[i], fo[i] = math.MaxInt32, math.MaxInt32, math.MaxInt32
	}
	id := func(x int) int {
		return x + off
	}
	fa[id(0)] = 0
	pre, z, o, res := 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		if s[i-1] == '1' {
			pre++
			o++
		} else {
			pre--
			z++
		}
		if j := fa[id(pre)]; j != math.MaxInt32 {
			res = max(res, i-j)
		}

		tar := id(pre - 2)
		if z < t0 {
			if j := fa[tar]; j != math.MaxInt32 {
				res = max(res, i-j)
			}
		} else {
			if j := fz[tar]; j != math.MaxInt32 {
				res = max(res, i-j)
			}
		}

		tar = id(pre + 2)
		if o < t1 {
			if j := fa[tar]; j != math.MaxInt32 {
				res = max(res, i-j)
			}
		} else {
			if j := fo[tar]; j != math.MaxInt32 {
				res = max(res, i-j)
			}
		}

		cur := id(pre)
		if fa[cur] == math.MaxInt32 {
			fa[cur] = i
		}
		if z > 0 && fz[cur] == math.MaxInt32 {
			fz[cur] = i
		}
		if o > 0 && fo[cur] == math.MaxInt32 {
			fo[cur] = i
		}
	}
	return res
}
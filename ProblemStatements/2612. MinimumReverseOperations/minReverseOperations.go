// func minReverseOperations(n int, p int, banned []int, k int) []int {
//     m := make(map[int]bool)
//     for _, v := range banned {
//         m[v] = true
//     }
//     a := make([]int, n)
//     for i := range a {
//         a[i] = -1
//     }
//     if m[p] {
//         return a
//     }
//     h := make(map[string]bool)
//     type s struct {
//         x []int
//         y int
//     }
//     i := make([]int, n)
//     i[p] = 1
//     q := []s{{i, 0}}
//     f := func(x []int) string {
//         r := make([]byte, n)
//         for i, v := range x {
//             if v == 1 {
//                 r[i] = '1'
//             } else {
//                 r[i] = '0'
//             }
//         }
//         return string(r)
//     }
//     g := func(x []int, t int) []int {
//         w := make([]int, n)
//         copy(w, x)
//         for i, j := t, t+k-1; i < j; i, j = i+1, j-1 {
//             w[i], w[j] = w[j], w[i]
//         }
//         return w
//     }
//     for len(q) > 0 {
//         c := q[0]
//         q = q[1:]
//         o := -1
//         for i, v := range c.x {
//             if v == 1 {
//                 o = i
//                 break
//             }
//         }
//         if a[o] == -1 {
//             a[o] = c.y
//         }
//         for i := 0; i <= n-k; i++ {
//             w := g(c.x, i)
//             u := -1
//             for j, v := range w {
//                 if v == 1 {
//                     u = j
//                     break
//                 }
//             }
//             if m[u] {
//                 continue
//             }
//             e := f(w)
//             if !h[e] {
//                 h[e] = true
//                 q = append(q, s{w, c.y + 1})
//             }
//         }
//     }
//     return a
// }

func minReverseOperations(n int, p int, banned []int, k int) []int {
	m := 1
	neven := n/2 + n&1
	for m < neven {
		m *= 2
	}
	var segtree [2][]bool
	segtree[0] = make([]bool, m*2)
	segtree[1] = make([]bool, m*2)

	var res [2][]int
	res[0] = make([]int, neven)
	res[1] = make([]int, n/2)
	for i := range res {
		for j := range res[i] {
			res[i][j] = -2
		}
	}
	mark := func(segtree []bool, res []int, i, val int) {
		segtree[m+i] = true
		for k := (m + i) / 2; k >= 1; k /= 2 {
			segtree[k] = segtree[k*2] && segtree[k*2+1]
			if !segtree[k] {
				break
			}
		}
		if i < len(res) {
			res[i] = val
		}
	}
	for _, i := range banned {
		mark(segtree[i&1], res[i&1], i/2, -1)
	}
	for i := neven; i < m; i++ {
		mark(segtree[0], res[0], i, -1)
	}
	for i := n / 2; i < m; i++ {
		mark(segtree[1], res[1], i, -1)
	}
	type pos struct {
		i   int
		odd int
	}
	var update func(segtree []bool, res []int, next *[]pos, i, lo, hi, qlo, qhi, val, odd int)
	update = func(segtree []bool, res []int, next *[]pos, i, lo, hi, qlo, qhi, val, odd int) {
		if qhi < lo || qlo > hi {
			return
		}

		if lo >= qlo && hi <= qhi {
			if segtree[i] {
				return
			}
			for j := lo; j <= hi; j++ {
				if res[j] != -2 {
					continue
				}
				res[j] = val
				mark(segtree, res, j, val)
				*next = append(*next, pos{j, odd})
			}
			return
		}
		mid := lo + (hi-lo)/2
		update(segtree, res, next, i*2, lo, mid, qlo, qhi, val, odd)
		update(segtree, res, next, i*2+1, mid+1, hi, qlo, qhi, val, odd)
	}
	curr := []pos{}
	next := []pos{}
	update(segtree[p&1], res[p&1], &curr, 1, 0, m-1, p/2, p/2, 0, p&1)
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			ii := x.i*2 + x.odd
			l, r := reachableRange(n, ii, k)
			update(segtree[l&1], res[l&1], &next, 1, 0, m-1, l/2, r/2, steps, l&1)
		}
		curr, next = next, curr
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = res[i&1][i/2]
	}
	for i := range ret {
		if ret[i] == -2 {
			ret[i] = -1
		}
	}
	return ret
}

func reachableRange(n, i, k int) (int, int) {
	leftMost := i - k + 1
	rightMost := i + k - 1
	if leftMost < 0 {
		d := -leftMost
		leftMost += d * 2
	}
	if rightMost >= n {
		d := rightMost - n + 1
		rightMost -= d * 2
	}
	return leftMost, rightMost
}
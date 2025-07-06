func minStable(nums []int, maxC int) int {
	n := len(nums)
	st, h := build(nums)
	l, r := 0, n
	for l < r {
		m := (l + r) / 2
		a := m + 1
		if a > n || h1(nums, st, h, a, maxC) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func h1(a []int, st [][]int, h []int, p, q int) bool {
	n := len(a)
	it := make([][2]int, 0, n)
	for i := 0; i+p <= n; i++ {
		if r(st, h, i, i+p-1) >= 2 {
			it = append(it, [2]int{i, i + p - 1})
		}
	}
	u := 0
	temp := -1
	for _, i := range it {
		if i[0] <= temp {
			continue
		}
		u++
		temp = i[1]
		if u > q {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func build(a []int) (st [][]int, l []int) {
	n := len(a)
	k := int(math.Floor(math.Log2(float64(n))) + 1)
	st = make([][]int, k)
	st[0] = make([]int, n)
	copy(st[0], a)
	for j := 1; j < k; j++ {
		st[j] = make([]int, n-(1<<j)+1)
		for i := 0; i+(1<<j) <= n; i++ {
			st[j][i] = gcd(st[j-1][i], st[j-1][i+(1<<(j-1))])
		}
	}
	l = make([]int, n+1)
	for i := 2; i <= n; i++ {
		l[i] = l[i>>1] + 1
	}
	return
}

func r(st [][]int, h []int, l, r int) int {
	j := h[r-l+1]
	return gcd(st[j][l], st[j][r-(1<<j)+1])
}
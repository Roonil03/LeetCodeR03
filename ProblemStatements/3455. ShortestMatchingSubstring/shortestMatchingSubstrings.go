func shortestMatchingSubstring(s string, p string) int {
	r := strings.Split(p, "*")
	if len(r) != 3 {
		return -1
	}
	a, b, c := r[0], r[1], r[2]
	n := len(s)
	if len(a)+len(b)+len(c) > n {
		return -1
	}
	A := o(s, a)
	C := o(s, c)
	if len(A) == 0 || len(C) == 0 {
		return -1
	}
	var N []int
	if b != "" {
		B := o(s, b)
		if len(B) == 0 {
			return -1
		}
		N = nO(n, B)
	}
	M := math.MaxInt32
	for _, x := range A {
		y := x + len(a)
		L := LB(C, y)
		if L == -1 {
			continue
		}
		u := y
		if b != "" {
			z := NB(N, y)
			if z == -1 {
				continue
			}
			u = max(u, z+len(b))
		}
		I := LB(C, u)
		if I == -1 {
			continue
		}
		d := C[I]
		if b != "" {
			w := N[y]
			if w == n || w+len(b) > d {
				v := false
				for j := I; j < len(C); j++ {
					if N[y] != n && N[y]+len(b) <= C[j] {
						d = C[j]
						v = true
						break
					}
				}
				if !v {
					continue
				}
			}
		}
		L2 := d + len(c) - x
		if L2 < M {
			M = L2
		}
	}
	if M == math.MaxInt32 {
		return -1
	}
	return M
}

func o(s, t string) []int {
	var r []int
	n := len(s)
	m := len(t)
	if m == 0 {
		for i := 0; i <= n; i++ {
			r = append(r, i)
		}
		return r
	}
	i := 0
	for {
		x := strings.Index(s[i:], t)
		if x == -1 {
			break
		}
		r = append(r, i+x)
		i = i + x + 1
		if i > n-m {
			break
		}
	}
	return r
}

func nO(n int, L []int) []int {
	N := make([]int, n+1)
	N[n] = n
	j := len(L) - 1
	for i := n - 1; i >= 0; i-- {
		if j >= 0 && L[j] == i {
			N[i] = L[j]
			j--
		} else {
			N[i] = N[i+1]
		}
	}
	return N
}

func LB(L []int, T int) int {
	l, r := 0, len(L)
	for l < r {
		mid := (l + r) / 2
		if L[mid] < T {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l < len(L) {
		return l
	}
	return -1
}

func NB(N []int, Y int) int {
	if Y < len(N) {
		if N[Y] < len(N)-1 {
			return N[Y]
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
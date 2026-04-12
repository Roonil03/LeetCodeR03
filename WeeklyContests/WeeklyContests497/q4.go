func countGoodSubseq(nums []int, p int, queries [][]int) int {
	n := len(nums)
	a := make([]int, n)
	dc := 0
	mx := 0
	for _, v := range nums {
		mx = max(mx, v)
	}
	for _, v := range queries {
		mx = max(mx, v[1])
	}
	for i, v := range nums {
		if v%p == 0 {
			dc++
			a[i] = v / p
		} else {
			a[i] = 0
		}
	}
	st := NewSegmentTree(a)
	// red := 0
	N := h1(mx/p) + 1
	// N := 7
	res := 0
	for _, q := range queries {
		id, v := q[0], q[1]
		if nums[id]%p == 0 {
			dc--
		}
		nums[id] = v
		if v%p == 0 {
			dc++
			a[id] = v / p
		} else {
			a[id] = 0
		}
		st.Update(id, a[id])
		g := st.a1()
		fg := false
		if dc > 0 {
			if dc < n {
				fg = (g == 1)
			} else {
				if g == 1 {
					if n > N {
						fg = true
					} else {
						fg = h2(a)
					}
				}
			}
		}
		if fg {
			res++
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func h1(lim int) int {
	if lim < 2 {
		return 0
	}
	primes := make([]int, 0, 16)
	prod, count := 1, 0
	for i := 2; ; i++ {
		fg := true
		for _, v := range primes {
			if v*v > i {
				break
			}
			if i%v == 0 {
				fg = false
				break
			}
		}
		if !fg {
			continue
		}
		if prod > lim/i {
			break
		}
		primes = append(primes, i)
		prod *= i
		count++
	}
	return count
}

func h2(a []int) bool {
	n := len(a)
	if n <= 1 {
		return false
	}
	pre := make([]int, n+1)
	suf := make([]int, n+1)
	for i := range n {
		pre[i+1] = gcd(pre[i], a[i])
	}
	for i := n - 1; i >= 0; i-- {
		suf[i] = gcd(suf[i+1], a[i])
	}
	for i := range n {
		if gcd(pre[i], suf[i+1]) == 1 {
			return true
		}
	}
	return false
}

type SegmentTree struct {
	n    int
	data []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]int, 2*size)
	for i := 0; i < n; i++ {
		data[size+i] = arr[i]
	}
	for i := size - 1; i > 0; i-- {
		data[i] = gcd(data[i<<1], data[2*i+1])
	}
	return &SegmentTree{n: size, data: data}
}

func (st *SegmentTree) Update(i, v int) {
	pos := i + st.n
	st.data[pos] = v
	for pos > 1 {
		pos >>= 1
		st.data[pos] = gcd(st.data[2*pos], st.data[2*pos+1])
	}
}

func (st *SegmentTree) a1() int {
	return st.data[1]
}
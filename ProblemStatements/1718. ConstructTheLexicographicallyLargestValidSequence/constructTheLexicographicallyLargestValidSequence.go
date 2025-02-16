func constructDistancedSequence(n int) []int {
	r := make([]int, 2*n-1)
	u := make([]bool, n+1)
	bt(r, u, n, 0)
	return r
}

func bt(r []int, u []bool, n, i int) bool {
	for i < len(r) && r[i] != 0 {
		i++
	}
	if i == len(r) {
		return true
	}
	for j := n; j >= 1; j-- {
		if u[j] {
			continue
		}
		if j == 1 {
			r[i] = 1
			u[1] = true
			if bt(r, u, n, i+1) {
				return true
			}
			r[i] = 0
			u[1] = false
		} else {
			if i+j < len(r) && r[i+j] == 0 {
				r[i] = j
				r[i+j] = j
				u[j] = true
				if bt(r, u, n, i+1) {
					return true
				}
				r[i] = 0
				r[i+j] = 0
				u[j] = false
			}
		}
	}
	return false
}
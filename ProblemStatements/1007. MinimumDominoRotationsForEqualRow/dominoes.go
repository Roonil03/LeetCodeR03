func minDominoRotations(tops []int, bottoms []int) int {
	a, b, s := make([]int, 7), make([]int, 7), make([]int, 7)
	n := len(tops)
	for i := 0; i < n; i++ {
		q := tops[i]
		p := bottoms[i]
		a[q]++
		b[p]++
		if q == p {
			s[q]++
		}
	}
	res := n
	for v := 1; v <= 6; v++ {
		if a[v]+b[v]-s[v] == n {
			minSwap := min(a[v], b[v]) - s[v]
			res = min(res, minSwap)
		}
	}
	if res == n {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
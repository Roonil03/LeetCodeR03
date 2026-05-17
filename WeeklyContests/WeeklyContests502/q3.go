func countLocalMaximums(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	a := m + 1
	b := (n + 1) * a
	pre := make([]int, 201*b)
	for v := 0; v <= 200; v++ {
		vs := v * b
		for i := range n {
			off1 := (i + 1) * a
			off0 := i * a
			for j := range m {
				var gt int = 0
				if matrix[i][j] > v {
					gt = 1
				}
				pre[vs+off1+j+1] = pre[vs+off0+j+1] + pre[vs+off1+j] - pre[vs+off0+j] + gt
			}
		}
	}
	res := 0
	for i := range n {
		for j := range m {
			x := matrix[i][j]
			if x == 0 {
				continue
			}
			r1, c1 := i-x, j-x
			r2, c2 := i+x, j+x
			ar1, ac1 := max(0, r1), max(0, c1)
			ar2, ac2 := min(n-1, r2), min(m-1, c2)
			vs := x * b
			tot := pre[vs+(ar2+1)*a+ac2+1] - pre[vs+ar1*a+ac2+1] - pre[vs+(ar2+1)*a+ac1] + pre[vs+ar1*a+ac1]
			if r1 >= 0 && c1 >= 0 && matrix[r1][c1] > x {
				tot--
			}
			if r1 >= 0 && c2 < m && matrix[r1][c2] > x {
				tot--
			}
			if r2 < n && c1 >= 0 && matrix[r2][c1] > x {
				tot--
			}
			if r2 < n && c2 < m && matrix[r2][c2] > x {
				tot--
			}
			if tot == 0 {
				res++
			}
		}
	}
	return res
}
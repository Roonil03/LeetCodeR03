func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	a := make([]bool, n)
	b := make([]bool, n)
	c := 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		a[A[i]-1] = true
		b[B[i]-1] = true
		if A[i] == B[i] {
			c++
		} else {
			if b[A[i]-1] {
				c++
			}
			if a[B[i]-1] {
				c++
			}
		}
		res[i] = c
	}
	return res
}
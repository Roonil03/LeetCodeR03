func kthCharacter(k int64, operations []int) byte {
	// type op struct{
	//     o, plen int
	// }
	n := len(operations)
	l := make([]int64, n+1)
	l[0] = 1
	for i := 0; i < n; i++ {
		if operations[i] == 0 {
			l[i+1] = l[i] * int64(2)
		} else {
			l[i+1] = l[i] * int64(2)
		}
		if l[i+1] > k {
			l[i+1] = k
		}
	}
	p := k - 1
	res := 0
	for i := n - 1; i >= 0; i-- {
		h := l[i]
		if p >= h {
			p -= h
			if operations[i] == 1 {
				res++
			}
		}
	}
	return byte('a' + (res % 26))
}
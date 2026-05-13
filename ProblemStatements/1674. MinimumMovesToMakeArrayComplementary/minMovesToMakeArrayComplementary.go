func minMoves(nums []int, limit int) int {
	n := len(nums)
	d := make([]int, 2*limit+2)
	for i := range n / 2 {
		u, v := nums[i], nums[n-1-i]
		if u > v {
			u, v = v, u
		}
		d[u+1]--
		d[u+v]--
		d[u+v+1]++
		d[v+limit+1]++
	}
	res, c := n, n
	for i := 2; i <= 2*limit; i++ {
		c += d[i]
		if c < res {
			res = c
		}
	}
	return res
}
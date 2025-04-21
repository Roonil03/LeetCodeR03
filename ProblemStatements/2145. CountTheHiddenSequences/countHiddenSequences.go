func numberOfArrays(differences []int, lower int, upper int) int {
	sum, n, x := 0, 0, 0
	for _, i := range differences {
		sum += i
		n = min(n, sum)
		x = max(x, sum)
	}
	return max(0, n-x+upper-lower+1)
}
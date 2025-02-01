func candy(ratings []int) int {
	n := len(ratings)
	c, i := n, 1
	for i < n {
		if ratings[i] == ratings[i-1] {
			i++
			continue
		}
		peak := 0
		for ratings[i] > ratings[i-1] {
			peak++
			c += peak
			i++
			if i == n {
				return c
			}
		}
		val := 0
		for i < n && ratings[i] < ratings[i-1] {
			val++
			c += val
			i++
		}
		c -= min(peak, val)
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
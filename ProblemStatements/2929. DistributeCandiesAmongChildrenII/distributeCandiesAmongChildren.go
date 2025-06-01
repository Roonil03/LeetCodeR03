func distributeCandies(n int, limit int) int64 {
	if n > 3*limit {
		return 0
	}
	if n == 0 {
		return 1
	}
	var total int64
	maxFirst := min(n, limit)
	for i := 0; i <= maxFirst; i++ {
		remaining := n - i
		maxSecond := min(remaining, limit)
		if remaining-maxSecond > limit {
			continue
		}
		minSecond := max(0, remaining-limit)
		if minSecond <= maxSecond {
			total += int64(maxSecond - minSecond + 1)
		}
	}
	return total
}
func min(a, b, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

func findMinimumOperations(s1, s2, s3 string) int {
	i := 0
	sz := min(len(s1), len(s2), len(s3))
	for i < sz && s1[i] == s2[i] && s2[i] == s3[i] {
		i++
	}
	if i == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - i*3
}

func partitionLabels(s string) []int {
	last := make(map[rune]int)

	for i, c := range s {
		last[c] = i
	}

	var result []int
	maxLast, prevEnd := 0, -1

	for i, c := range s {
		if last[c] > maxLast {
			maxLast = last[c]
		}

		if i == maxLast {
			result = append(result, i-prevEnd)
			prevEnd = i
		}
	}

	return result
}
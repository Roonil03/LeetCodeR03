func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, n := range arr {
		count[n]++
	}
	seen := make(map[int]struct{})
	for _, f := range count {
		if _, ex := seen[f]; ex {
			return false
		}
		seen[f] = struct{}{}
	}
	return true
}
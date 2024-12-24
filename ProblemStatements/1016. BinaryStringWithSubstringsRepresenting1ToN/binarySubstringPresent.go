func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		binRep := strconv.FormatInt(int64(i), 2)
		if !strings.Contains(s, binRep) {
			return false
		}
	}
	return true
}
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}
	mp := make(map[string]bool)
	for i := 0; i <= len(s)-k; i++ {
		mp[s[i:i+k]] = true
	}
	return len(mp) == 1<<k
}
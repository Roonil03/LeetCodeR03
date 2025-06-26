func longestSubsequence(s string, k int) int {
	n, res, val, pow := len(s), 0, int64(0), 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			res++
		} else {
			if pow < 32 {
				val += 1 << pow
				if val <= int64(k) {
					res++
				}
			}
		}
		pow++
	}
	return res
}
func smallestPalindrome(s string) string {
	temp := make([]int, 26)
	for _, ch := range s {
		temp[ch-'a']++
	}
	res := make([]byte, len(s))
	l, r := 0, len(res)-1
	mid := byte(0)
	for i := 0; i < 26; i++ {
		for temp[i] >= 2 {
			b := byte('a' + i)
			res[l], res[r] = b, b
			l++
			r--
			temp[i] -= 2
		}
		if temp[i] == 1 {
			mid = byte('a' + i)
		}
	}
	if l == r {
		res[l] = mid
	}
	return string(res)
}
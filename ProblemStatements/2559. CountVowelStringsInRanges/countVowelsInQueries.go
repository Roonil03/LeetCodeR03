func vowelStrings(words []string, queries [][]int) []int {
	n1 := len(words)
	n2 := len(queries)
	q := make([]int, n1)
	res := make([]int, n2)
	count := 0
	for i := 0; i < n1; i++ {
		if isVowel(words[i][0], words[i][len(words[i])-1]) {
			count++
		}
		q[i] = count
	}
	for i := 0; i < n2; i++ {
		li, ri := queries[i][0], queries[i][1]
		if li == 0 {
			res[i] = q[ri]
		} else {
			res[i] = q[ri] - q[li-1]
		}
	}
	return res
}

func isVowel(c1, c2 byte) bool {
	vowels := "aeiou"
	return contains(vowels, c1) && contains(vowels, c2)
}

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
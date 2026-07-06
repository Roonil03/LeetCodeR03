func minOperations(s1 string, s2 string) int {
	if len(s1) == 1 {
		if s1 == s2 {
			return 0
		}
		if s1 == "0" && s2 == "1" {
			return 1
		}
		return -1
	}
	a, b, p := 0, 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == '0' && s2[i] == '1' {
			b++
		} else if s1[i] == '1' && s2[i] == '0' {
			a++
			if i < len(s1)-1 && s1[i+1] == '1' && s2[i+1] == '0' {
				p++
				a++
				i++
			}
		}
	}
	return b + a*2 - p*3
}
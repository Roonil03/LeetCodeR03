func maximumGain(s string, x int, y int) int {
	if len(s) < 2 {
		return 0
	}
	var a, b string
	var a1, b1 int
	if x >= y {
		a, a1 = "ab", x
		b, b1 = "ba", y
	} else {
		a, a1 = "ba", y
		b, b1 = "ab", x
	}
	tot, res := 0, []byte{}
	temp := []byte{}
	for i := 0; i < len(s); i++ {
		if len(temp) > 0 && temp[len(temp)-1] == a[0] && s[i] == a[1] {
			temp = temp[:len(temp)-1]
			tot += a1
		} else {
			temp = append(temp, s[i])
		}
	}
	for i := 0; i < len(temp); i++ {
		if len(res) > 0 && res[len(res)-1] == b[0] && temp[i] == b[1] {
			res = res[:len(res)-1]
			tot += b1
		} else {
			res = append(res, temp[i])
		}
	}
	return tot
}
func minOperations(s string) int {
	a, b := 0, 0
	a1, b1 := true, false
	// res := 0
	for i := range s {
		if a1 && s[i] != '1' {
			a++
		} else if !a1 && s[i] != '0' {
			a++
		}
		if b1 && s[i] != '1' {
			b++
		} else if !b1 && s[i] != '0' {
			b++
		}
		a1 = !a1
		b1 = !b1
	}
	return min(a, b)
	//     res += (int(s[i]) ^ i) & 1;
	// }
	// return min(res, len(s) - res)
}
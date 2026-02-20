func makeLargestSpecial(s string) string {
	if len(s) < 1 {
		return ""
	}
	var bal, start int
	var res []string
	for j := 0; j < len(s); j++ {
		if s[j] == '1' {
			bal++
		} else {
			bal--
		}
		if bal == 0 {
			res = append(res, fmt.Sprintf("1%s0", makeLargestSpecial(s[start+1:j])))
			start = j + 1
		}
	}
	if len(res) < 1 {
		return ""
	}
	if len(res) == 1 {
		return res[0]
	}
	sort.Slice(res, func(i, j int) bool {
		return res[j] < res[i]
	})
	return strings.Join(res, "")
}
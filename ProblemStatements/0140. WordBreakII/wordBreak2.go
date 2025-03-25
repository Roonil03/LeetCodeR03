func wordBreak(s string, wordDict []string) []string {
	set := make(map[string]bool)
	for _, i := range wordDict {
		set[i] = true
	}
	res := []string{}
	helper(0, s, "", set, &res)
	return res
}

func helper(id int, s, cur string, set map[string]bool, res *[]string) {
	if id == len(s) {
		*res = append(*res, cur[:len(cur)-1])
		return
	}
	str := ""
	for i := id; i < len(s); i++ {
		str += string(s[i])
		if set[str] {
			helper(i+1, s, cur+str+" ", set, res)
		}
	}
}
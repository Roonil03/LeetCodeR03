func longestBalanced(s string) int {
	n := len(s)
	res := 0
	for i := range s {
		mp := [26]int{}
		for j := i; j < n; j++ {
			mp[int(s[j]-'a')]++
			dist, f := 0, 0
			for k := range mp {
				if mp[k] > 0 {
					dist++
					if mp[k] > f {
						f = mp[k]
					}
				}
			}
			if dist > 0 && dist*f == j-i+1 {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}
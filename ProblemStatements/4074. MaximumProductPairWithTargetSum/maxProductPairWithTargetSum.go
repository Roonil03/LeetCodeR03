func maxProductPair(nums []int, target int) []int {
	res := []int{-1, -1}
	a, s := 0, 2-(target&1)
	mp := make(map[int]int)
	for i, v := range nums {
		u := target - v
		if j, ok := mp[u]; ok && u != v {
			d := u - v
			p0, p1 := j, i
			if d < 0 {
				d, p0, p1 = -d, i, j
			}
			if res[0] == -1 || d < a {
				a, res[0], res[1] = d, p0, p1
				if d == s {
					return res
				}
			}
		}
		mp[v] = i
	}
	return res
}
func totalFruit(fruits []int) int {
	n := len(fruits)
	l, res := 0, 0
	count := make(map[int]int)
	for r := 0; r < n; r++ {
		count[fruits[r]]++
		for len(count) > 2 {
			count[fruits[l]]--
			if count[fruits[l]] == 0 {
				delete(count, fruits[l])
			}
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
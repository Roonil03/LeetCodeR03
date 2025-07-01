func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n, m := len(spells), len(potions)
	res := make([]int, n)
	for i, spell := range spells {
		left, right := 0, m
		for left < right {
			mid := left + (right-left)/2
			if int64(spell)*int64(potions[mid]) >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		res[i] = m - left
	}
	return res
}
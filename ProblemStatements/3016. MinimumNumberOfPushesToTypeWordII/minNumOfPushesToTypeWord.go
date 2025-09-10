func minimumPushes(word string) int {
	count := make([]int, 26)
	for i := 0; i < len(word); i++ {
		count[int(word[i]-'a')]++
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})
	res := 0
	for i := 0; i < 26; i++ {
		if count[i] == 0 {
			break
		}
		cost := i/8 + 1
		res += cost * count[i]
	}
	return res
}
func subarrayBitwiseORs(arr []int) int {
	res := make(map[int]struct{})
	cur := map[int]struct{}{}
	for _, num := range arr {
		next := map[int]struct{}{num: {}}
		for x := range cur {
			next[x|num] = struct{}{}
		}
		for x := range next {
			res[x] = struct{}{}
		}
		cur = next
	}
	return len(res)
}
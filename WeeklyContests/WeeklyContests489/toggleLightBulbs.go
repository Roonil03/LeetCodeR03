func toggleLightBulbs(bulbs []int) []int {
	mp := make([]bool, 100)
	for i := range bulbs {
		mp[bulbs[i]-1] = !mp[bulbs[i]-1]
	}
	res := make([]int, 0)
	for i := range mp {
		if mp[i] {
			res = append(res, i+1)
		}
	}
	return res
}
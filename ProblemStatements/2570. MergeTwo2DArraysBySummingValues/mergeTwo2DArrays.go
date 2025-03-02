func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num[0]] += num[1]
	}
	for _, num := range nums2 {
		m[num[0]] += num[1]
	}
	res := make([][]int, 0, len(m))
	for id, val := range m {
		res = append(res, []int{id, val})
	}
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i][0] > res[j][0] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}
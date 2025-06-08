func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}
	var a, b []int
	for n := range m1 {
		if _, ok := m2[n]; !ok {
			a = append(a, n)
		}
	}
	for n := range m2 {
		if _, ok := m1[n]; !ok {
			b = append(b, n)
		}
	}
	return [][]int{a, b}
}
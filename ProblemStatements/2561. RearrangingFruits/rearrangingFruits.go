func minCost(basket1 []int, basket2 []int) int64 {
	count := make(map[int]int)
	a := basket1[0]
	for _, v := range basket1 {
		count[v]++
		if v < a {
			a = v
		}
	}
	for _, v := range basket2 {
		count[v]--
		if v < a {
			a = v
		}
	}
	nums := make([]int, 0)
	for i, j := range count {
		if j%2 != 0 {
			return int64(-1)
		}
		for k := 0; k < abs(j)/2; k++ {
			nums = append(nums, i)
		}
	}
	sort.Ints(nums)
	m := len(nums)
	res := int64(0)
	for i := 0; i < m/2; i++ {
		res += int64(min(nums[i], a*2))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
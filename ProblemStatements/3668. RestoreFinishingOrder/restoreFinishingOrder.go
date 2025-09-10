func recoverOrder(order []int, friends []int) []int {
	res := make([]int, len(friends))
	k := 0
	for _, i := range order {
		for _, j := range friends {
			if i == j {
				res[k] = i
				k++
			}
		}
	}
	return res
}
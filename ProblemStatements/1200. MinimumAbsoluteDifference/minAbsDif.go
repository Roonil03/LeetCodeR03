func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	m := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) < m {
			m = abs(arr[i+1] - arr[i])
		}
	}
	res := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) == m {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
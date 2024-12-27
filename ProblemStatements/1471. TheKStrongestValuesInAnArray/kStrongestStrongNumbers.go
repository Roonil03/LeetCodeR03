func abs(x float64) int {
	if x < 0 {
		return int(-x)
	}
	return int(x)
}
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	n := len(arr)
	med := arr[(n-1)/2]
	sort.Slice(arr, func(i, j int) bool {
		dif1 := abs(float64(arr[i] - med))
		dif2 := abs(float64(arr[j] - med))
		if dif1 == dif2 {
			return arr[i] > arr[j]
		}
		return dif1 > dif2
	})
	return arr[:k]
}
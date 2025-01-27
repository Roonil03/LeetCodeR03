func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	ex, a, s := 0, 0, 0
	for i := 0; i < n; i++ {
		ex += gas[i] - cost[i]
		a += gas[i] - cost[i]
		if a < 0 {
			a = 0
			s = i + 1
		}
	}
	if ex < 0 {
		return -1
	}
	return s
}
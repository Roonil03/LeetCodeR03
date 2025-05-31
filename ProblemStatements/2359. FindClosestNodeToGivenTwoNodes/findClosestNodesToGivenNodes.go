func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	m1 := make([]int, n)
	m2 := make([]int, n)
	for i := 0; i < n; i++ {
		m1[i] = -1
		m2[i] = -1
	}
	fillDist := func(start int, memo []int) {
		dist, i := 0, start
		for i != -1 && memo[i] == -1 {
			memo[i] = dist
			dist++
			i = edges[i]
		}
	}
	fillDist(node1, m1)
	fillDist(node2, m2)
	res, minDist := -1, int(1e9)
	for i := 0; i < n; i++ {
		if m1[i] >= 0 && m2[i] >= 0 {
			d := m1[i]
			if m2[i] > d {
				d = m2[i]
			}
			if d < minDist || (d == minDist && i < res) {
				minDist = d
				res = i
			}
		}
	}
	return res
}
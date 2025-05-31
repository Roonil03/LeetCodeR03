func dfs(i, p int, al [][]int, parity []bool, even bool) int {
	targets := 0
	if even {
		targets = 1
	}
	parity[i] = even
	for _, j := range al[i] {
		if j != p {
			targets += dfs(j, i, al, parity, !even)
		}
	}
	return targets
}

func adjacencyList(edges [][]int) [][]int {
	al := make([][]int, len(edges)+1)
	for _, e := range edges {
		al[e[0]] = append(al[e[0]], e[1])
		al[e[1]] = append(al[e[1]], e[0])
	}
	return al
}

func maxTargetNodes(e1, e2 [][]int) []int {
	m, n := len(e1)+1, len(e2)+1
	parity := make([]bool, m)
	ignored := make([]bool, n)
	even1 := dfs(0, -1, adjacencyList(e1), parity, true)
	odd1 := m - even1
	even2 := dfs(0, -1, adjacencyList(e2), ignored, true)
	odd2 := n - even2
	res := make([]int, m)
	for i := 0; i < m; i++ {
		val := max(even2, odd2)
		if parity[i] {
			val += even1
		} else {
			val += odd1
		}
		res[i] = val
	}
	return res
}
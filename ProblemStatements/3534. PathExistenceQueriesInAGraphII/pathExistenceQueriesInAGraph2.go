func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	pos := make([]int, n)
	for i := range pos {
		pos[i] = i
	}
	slices.SortFunc(pos, func(i, j int) int {
		return nums[i] - nums[j]
	})
	rev := make([]int, n)
	for i, v := range pos {
		rev[v] = i
	}
	comp := make([]int, n)
	far := make([]int, n)
	r := 0
	for i := range n {
		if i > 0 {
			if nums[pos[i]]-nums[pos[i-1]] > maxDiff {
				comp[i] = comp[i-1] + 1
			} else {
				comp[i] = comp[i-1]
			}
		}
		for r < n && nums[pos[r]]-nums[pos[i]] <= maxDiff {
			r++
		}
		far[i] = r - 1
	}
	const bits = 18
	up := make([][bits]int, n)
	for i := range n {
		up[i][0] = far[i]
	}
	for i := 1; i < bits; i++ {
		for j := range n {
			up[j][i] = up[up[j][i-1]][i-1]
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		u, v := rev[q[0]], rev[q[1]]
		if u == v {
			res[i] = 0
			continue
		}
		if u > v {
			u, v = v, u
		}
		if comp[u] != comp[v] {
			res[i] = -1
			continue
		}
		temp := 0
		for j := bits - 1; j >= 0; j-- {
			if up[u][j] < v {
				u = up[u][j]
				temp += 1 << j
			}
		}
		res[i] = temp + 1
	}
	return res
}
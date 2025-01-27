func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)
	for i := range adj {
		adj[i] = make([]int, numCourses)
	}
	for _, pre := range prerequisites {
		adj[pre[0]][pre[1]] = 1
	}
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if adj[i][k] == 1 && adj[k][j] == 1 {
					adj[i][j] = 1
				}
			}
		}
	}
	res := make([]bool, len(queries))
	for i, query := range queries {
		res[i] = adj[query[0]][query[1]] == 1
	}
	return res
}
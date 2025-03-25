func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		g[i] = make([]int, 0)
	}
	for _, i := range prerequisites {
		c := i[0]
		p := i[1]
		g[c] = append(g[c], p)
	}
	vis := make([]int, numCourses)
	for c := 0; c < numCourses; c++ {
		if vis[c] == 0 {
			if hasCycle(c, g, vis) {
				return false
			}
		}
	}
	return true
}

func hasCycle(c int, g [][]int, vis []int) bool {
	if vis[c] == 1 {
		return true
	}
	if vis[c] == 2 {
		return false
	}
	vis[c] = 1
	for _, prereq := range g[c] {
		if hasCycle(prereq, g, vis) {
			return true
		}
	}
	vis[c] = 2
	return false
}
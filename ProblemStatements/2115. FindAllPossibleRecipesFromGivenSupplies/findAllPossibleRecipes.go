func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Build graph and in-degree map
	//n := len(recipes)
	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	recipeSet := make(map[string]bool)
	// Initialize recipe set for O(1) lookup
	for _, recipe := range recipes {
		recipeSet[recipe] = true
		inDegree[recipe] = 0
	}
	// Build graph and count dependencies
	for i, recipe := range recipes {
		for _, ing := range ingredients[i] {
			if recipeSet[ing] {
				graph[ing] = append(graph[ing], recipe)
				inDegree[recipe]++
			}
		}
	}
	// Initialize queue with available supplies
	queue := []string{}
	available := make(map[string]bool)
	for _, supply := range supplies {
		available[supply] = true
	}
	// Add recipes with all ingredients available
	for i, recipe := range recipes {
		canMake := true
		for _, ing := range ingredients[i] {
			if !available[ing] && !recipeSet[ing] {
				canMake = false
				break
			}
		}
		if canMake && inDegree[recipe] == 0 {
			queue = append(queue, recipe)
		}
	}
	// Process queue (topological sort)
	result := []string{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if recipeSet[curr] {
			result = append(result, curr)
			available[curr] = true
		}
		// Process neighbors
		for _, next := range graph[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				canMake := true
				// Check if all ingredients are available
				for i, recipe := range recipes {
					if recipe == next {
						for _, ing := range ingredients[i] {
							if !available[ing] && !recipeSet[ing] {
								canMake = false
								break
							}
						}
						break
					}
				}
				if canMake {
					queue = append(queue, next)
				}
			}
		}
	}
	return result
}
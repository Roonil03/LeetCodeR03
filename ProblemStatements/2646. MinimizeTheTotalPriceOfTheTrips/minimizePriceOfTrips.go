// func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
//     adj := make([][]int, n)
//     for _, i:= range edges{
//         adj[i[0]] = append(adj[i[0]], i[1])
//         adj[i[1]] = append(adj[i[1]], i[0])
//     }
//     freq := make([]int, n)
//     for _, i := range trips{
//         path := findPath(adj, i[0], i[1], n)
//         if len(path) >0{
//             for _, p := range path{
//                 freq[p.node]++
//             }
//         }
//     }
//     for i:= range price{
//         price[i] *= freq[i]
//     }
//     normal, halved := dp(0, -1, adj, price)
//     return min(normal, halved)
// }

// type Path struct {
// 	node   int
// 	parent int
// 	dist   int
// }

// func findPath(adj [][]int, s, e, n int)[]Path{
//     vis := make([]bool, n)
//     path := []Path{}
//     found := false
//     var dfs func(node, parent, dist int)
// 	dfs = func(node, parent, dist int) {
// 		if found {
// 			return
// 		}
// 		vis[node] = true
// 		path = append(path, Path{node, parent, dist})
// 		if node == e {
// 			found = true
// 			return
// 		}
// 		for _, next := range adj[node] {
// 			if !vis[next] {
// 				dfs(next, node, dist+1)
// 			}
// 		}
// 		if !found {
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	dfs(s, -1, 0)
// 	if !found {
// 		return nil
// 	}
// 	return path
// }

// func dp(node, parent int, adj [][]int, price []int) (int, int) {
//     totalNormal := price[node]
//     totalHalved := price[node] / 2
//     for _, next := range adj[node] {
//         if next != parent {
//             nextNormal, nextHalved := dp(next, node, adj, price)
//             totalNormal += min(nextNormal, nextHalved)
//             totalHalved += nextNormal
//         }
//     }
//     return totalNormal, totalHalved
// }

// func min(a, b int) int{
//     if a > b{
//         return b
//     }
//     return a
// }

func dfsPathCount(graph [][]int, start int, end int, parent int, pathCount []int) bool {
	if start == end {
		pathCount[start]++
		return true
	}
	for _, neighbour := range graph[start] {
		if neighbour == parent {
			continue
		}
		if dfsPathCount(graph, neighbour, end, start, pathCount) {
			pathCount[start]++
			return true
		}
	}
	return false
}

func treeDp(graph [][]int, node int, parent int, pathCount []int, price []int) (int, int) {
	noDiscount := pathCount[node] * price[node]
	discount := pathCount[node] * price[node] / 2
	for _, neighbour := range graph[node] {
		if neighbour == parent {
			continue
		}
		childDiscount, childNoDiscount := treeDp(graph, neighbour, node, pathCount, price)
		discount += childNoDiscount
		noDiscount += min(childDiscount, childNoDiscount)
	}
	return discount, noDiscount
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u], adj[v] = append(adj[u], v), append(adj[v], u)
	}
	pathCount := make([]int, n)
	for _, trip := range trips {
		dfsPathCount(adj, trip[0], trip[1], -1, pathCount)
	}
	return min(treeDp(adj, 0, -1, pathCount, price))
}
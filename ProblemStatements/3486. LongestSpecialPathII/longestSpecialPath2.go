func longestSpecialPath(edges [][]int, nums []int) []int {
	//     n := len(nums)
	//     adj := make([][][2]int, n)
	//     for _, e := range edges {
	//         u, v, l := e[0], e[1], e[2]
	//         adj[u] = append(adj[u], [2]int{v, l})
	//         adj[v] = append(adj[v], [2]int{u, l})
	//     }
	//     parent := make([]int, n)
	//     for i := range parent {
	//         parent[i] = -2
	//     }
	//     parent[0] = -1
	//     queue := []int{0}
	//     for len(queue) > 0 {
	//         u := queue[0]
	//         queue = queue[1:]
	//         for _, p := range adj[u] {
	//             v := p[0]
	//             if parent[v] == -2 {
	//                 parent[v] = u
	//                 queue = append(queue, v)
	//             }
	//         }
	//     }
	//     res, mn := 0, math.MaxInt32
	//     for i := 0; i < n; i++ {
	//         vis := make([]bool, n)
	//         val := make(map[int]int)
	//         dfs(i, parent[i], adj, nums, vis, val, 0, 1, &res, &mn)
	//     }
	//     return []int{res, mn}
	// }

	// func dfs(node, par int, adj [][][2]int, nums []int, vis []bool, val map[int]int, l, n int, res, mn *int) {
	//     vis[node] = true
	//     val[nums[node]]++
	//     dup := 0
	//     for _, ch := range val {
	//         if ch > 2 {
	//             vis[node] = false
	//             val[nums[node]]--
	//             if val[nums[node]] == 0 {
	//                 delete(val, nums[node])
	//             }
	//             return
	//         }
	//         if ch == 2 {
	//             dup++
	//         }
	//     }
	//     if dup > 1 {
	//         vis[node] = false
	//         val[nums[node]]--
	//         if val[nums[node]] == 0 {
	//             delete(val, nums[node])
	//         }
	//         return
	//     }
	//     if l > *res {
	//         *res = l
	//         *mn = n
	//     } else if l == *res && n < *mn {
	//         *mn = n
	//     }
	//     for _, n1 := range adj[node] {
	//         nn, el := n1[0], n1[1]
	//         if nn != par && !vis[nn] {
	//             dfs(nn, node, adj, nums, vis, val, l+el, n+1, res, mn)
	//         }
	//     }
	//     vis[node] = false
	//     val[nums[node]]--
	//     if val[nums[node]] == 0 {
	//         delete(val, nums[node])
	//     }
	n := len(edges) + 1
	paths := make([][]int, n)
	mxNum := 0
	for i := range nums {
		if nums[i] > mxNum {
			mxNum = nums[i]
		}
	}
	for i := range edges {
		if paths[edges[i][0]] == nil {
			paths[edges[i][0]] = make([]int, 0)
		}
		if paths[edges[i][1]] == nil {
			paths[edges[i][1]] = make([]int, 0)
		}
		paths[edges[i][0]] = append(paths[edges[i][0]], i)
		paths[edges[i][1]] = append(paths[edges[i][1]], i)
	}
	maxSpecial := 0
	minNum := 0
	p := make([]int, n)
	pSum := make([]int, n+1)
	for _, i := range paths[0] {
		if edges[i][1] == 0 {
			edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
		}
		mapCheck := make([]int, mxNum+1)
		mapCheck[nums[edges[i][0]]] = 1
		pNext := make([]int, 0)
		if nums[edges[i][1]] == nums[edges[i][0]] {
			pNext = append(pNext, 1)
		}
		mapCheck[nums[edges[i][1]]] = 2
		p[0] = i
		pSum[1] = edges[i][2]
		rsi := getSpecialPath(0, p, pSum, pNext, nums, edges, paths, 1, mapCheck)
		if rsi[0] > maxSpecial {
			maxSpecial = rsi[0]
			minNum = rsi[1]
		} else if rsi[0] == maxSpecial && rsi[1] < minNum {
			minNum = rsi[1]
		}
	}
	return []int{maxSpecial, minNum}
}

func getSpecialPath(start int, p, pSum, pNext, nums []int, edges, paths [][]int, idx int, mapCheck []int) []int {
	rs := make([]int, 2)
	for len(pNext) >= 2 {
		if pNext[0] > pNext[1] {
			pNext[0], pNext[1] = pNext[1], pNext[0]
		}
		start = pNext[0]
		pNext = pNext[1:]
	}
	rs[0] = pSum[idx] - pSum[start]
	rs[1] = idx - start + 1
	last := edges[p[idx-1]][1]
	if len(paths[last]) > 1 {
		for _, i := range paths[last] {
			if i == p[idx-1] {
				continue
			}
			if last == edges[i][1] {
				edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
			}
			lastIndex := mapCheck[nums[edges[i][1]]]
			mapCheck[nums[edges[i][1]]] = idx + 2
			pNextNew := make([]int, len(pNext))
			copy(pNextNew, pNext)
			if lastIndex > start {
				pNextNew = append(pNextNew, lastIndex)
				if (lastIndex < idx || (lastIndex == idx && nums[edges[i][1]] == nums[edges[p[idx-1]][0]])) && nums[edges[i][1]] == nums[edges[i][0]] {
					pNextNew = append(pNextNew, idx+1)
					mapCheck[nums[edges[i][1]]] = idx + 2
				}
			}
			p[idx] = i
			pSum[idx+1] = pSum[idx] + edges[i][2]
			rsi := getSpecialPath(start, p, pSum, pNextNew, nums, edges, paths, idx+1, mapCheck)
			if rsi[0] > rs[0] {
				rs[0] = rsi[0]
				rs[1] = rsi[1]
			} else if rsi[0] == rs[0] && rsi[1] < rs[1] {
				rs[1] = rsi[1]
			}
			mapCheck[nums[edges[i][1]]] = lastIndex
			p[idx] = 0
			pSum[idx+1] = 0
		}
	}
	return rs
}
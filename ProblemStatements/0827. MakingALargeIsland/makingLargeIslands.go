// 931 ms Code:
// func largestIsland(grid [][]int) int {
//     res := 0
//     for i := 0; i < len(grid); i++ {
//         for j := 0; j < len(grid[i]); j++ {
//             if grid[i][j] == 0 {
//                 res = max(res, paint(i, j, 2, grid, true))
//                 paint(i, j, 1, grid, true)
//             }
//         }
//     }
//     if res == 0 {
//         return len(grid) * len(grid[0])
//     }
//     return res
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func paint(i, j, c int, grid [][]int, flip bool) int {
//     if !flip && (min(i, j) < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 || grid[i][j] == c) {
//         return 0
//     }
//     if grid[i][j] == 0 {
//         grid[i][j] = 0
//     } else {
//         grid[i][j] = c
//     }
//     return 1 + paint(i+1, j, c, grid, false) + paint(i-1, j, c, grid, false) + paint(i, j+1, c, grid, false) + paint(i, j-1, c, grid, false)
// }

// func min(a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }

// 137 ms Code:
// type UnionFind struct {
// 	parent []int
// 	size   []int
// }

// func newUnionFind(n int) *UnionFind {
// 	uf := &UnionFind{
// 		parent: make([]int, n),
// 		size:   make([]int, n),
// 	}
// 	for i := range uf.parent {
// 		uf.parent[i] = -1
// 	}
// 	return uf
// }

// func (uf *UnionFind) isExist(u int) bool {
// 	return uf.parent[u] >= 0
// }

// func (uf *UnionFind) add(u int) {
// 	if uf.isExist(u) {
// 		return
// 	}
// 	uf.parent[u] = u
// 	uf.size[u] = 1
// }

// func (uf *UnionFind) find(u int) int {
// 	if uf.parent[u] == u {
// 		return u
// 	}
// 	uf.parent[u] = uf.find(uf.parent[u])
// 	return uf.parent[u]
// }

// func (uf *UnionFind) union(u, v int) bool {
// 	pu, pv := uf.find(u), uf.find(v)
// 	if pu == pv {
// 		return false
// 	}
// 	if uf.size[pu] <= uf.size[pv] {
// 		uf.parent[pu] = pv
// 		uf.size[pv] += uf.size[pu]
// 	} else {
// 		uf.parent[pv] = pu
// 		uf.size[pu] += uf.size[pv]
// 	}
// 	return true
// }

// func largestIsland(grid [][]int) int {
// 	DIR := []int{0, 1, 0, -1, 0}
// 	m, n, ans := len(grid), len(grid[0]), 0
// 	uf := newUnionFind(m * n)
// 	landNeighbors := func(r, c int) []int {
// 		var neighbors []int
// 		for i := 0; i < 4; i++ {
// 			nr, nc := r+DIR[i], c+DIR[i+1]
// 			neiId := nr*n + nc
// 			if nr < 0 || nr >= m || nc < 0 || nc >= n || !uf.isExist(neiId) {
// 				continue
// 			}
// 			neighbors = append(neighbors, neiId)
// 		}
// 		return neighbors
// 	}
// 	for r := 0; r < m; r++ {
// 		for c := 0; c < n; c++ {
// 			if grid[r][c] == 0 {
// 				continue
// 			}
// 			curId := r*n + c
// 			uf.add(curId)
// 			for _, neiId := range landNeighbors(r, c) {
// 				uf.union(curId, neiId)
// 			}
// 			p := uf.find(curId)
// 			if uf.size[p] > ans {
// 				ans = uf.size[p]
// 			}
// 		}
// 	}

// 	for r := 0; r < m; r++ {
// 		for c := 0; c < n; c++ {
// 			if grid[r][c] == 1 {
// 				continue
// 			}
// 			neiParents := make(map[int]struct{})
// 			for _, neiId := range landNeighbors(r, c) {
// 				neiParents[uf.find(neiId)] = struct{}{}
// 			}
// 			sizeFormed := 1
// 			for p := range neiParents {
// 				sizeFormed += uf.size[p]
// 			}
// 			if sizeFormed > ans {
// 				ans = sizeFormed
// 			}
// 		}
// 	}
// 	return ans
// }

//88 ms Solution Code:
type DS struct {
	size           []int
	ultimateParent []int
}

func (d *DS) union(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	if pu == pv {
		return
	}
	if d.size[pu] >= d.size[pv] {
		d.ultimateParent[pv] = pu
		d.size[pu] += d.size[pv]
	} else {
		d.ultimateParent[pu] = pv
		d.size[pv] += d.size[pu]
	}
}

func (d *DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}

func createDS(nodes int) *DS {
	size := make([]int, nodes)
	up := make([]int, nodes)
	for i := 0; i < nodes; i++ {
		size[i] = 1
		up[i] = i
	}
	return &DS{size: size, ultimateParent: up}
}

func largestIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	ds := createDS(rows * cols)
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 1 {
				directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
				for _, direction := range directions {
					nx, ny := x+direction[0], y+direction[1]
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
						ds.union(nx*cols+ny, x*cols+y)
					}
				}
			}
		}
	}

	maxCount := -1
	zeroPresent := false
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 0 {
				zeroPresent = true
				count := 1
				parents := map[int]bool{}
				directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
				for _, direction := range directions {
					nx, ny := x+direction[0], y+direction[1]
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
						up := ds.findUltimateParent(nx*cols + ny)
						if !parents[up] {
							count += ds.size[up]
							parents[up] = true
						}
					}
				}
				if count > maxCount {
					maxCount = count
				}
			}
		}
	}
	if zeroPresent {
		return maxCount
	}
	return rows * cols
}

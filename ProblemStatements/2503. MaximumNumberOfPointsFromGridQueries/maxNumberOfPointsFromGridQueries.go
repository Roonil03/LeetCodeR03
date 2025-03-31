// func maxPoints(g [][]int, q []int) []int {
//     sort.Ints(queries)
//     res := make([]int, len(queries))
//     d := [][2]int{{0,1}, {1,0},{0, -1},{-1,0}}
//     m,n := len(grid), len(grid[0])
//     bfs := func(t int) int{
//         vis := make([][]bool, m)
//         for i:= range vis{
//             vis[i] = make([]bool, n)
//         }
//         q := [][3]int{{0,0,0}}
//         mps := 0
//         for len(q) > 0{
//             r, c, p := q[0][0], q[0][1], q[0][2]
//             q = q[1:]
//             if r < 0 || r >= m || c < 0 || c >= n || vis[r][c] || grid[r][c] >= t {
//                 continue
//             }
//             vis[r][c] = true
//             mps = p + 1
//             for _, dir := range d{
//                 nr, nc := r + dir[0], c + dir[1]
//                 q = append(q, [3]int{nr, nc, mps})
//             }
//         }
//         return mps
//     }
//     for i, t := range queries{
//         res[i] = bfs(t)
//     }
//     return res
// }

func maxPoints(grid [][]int, queries []int) []int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	sortedQueries := make([][]int, len(queries))
	for i, val := range queries {
		sortedQueries[i] = []int{val, i}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][0] < sortedQueries[j][0]
	})
	result := make([]int, len(queries))
	h := &MinHeap{}
	heap.Init(h)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	heap.Push(h, []int{grid[0][0], 0, 0})
	visited[0][0] = true
	points := 0
	for _, q := range sortedQueries {
		queryVal, queryIdx := q[0], q[1]
		for h.Len() > 0 && (*h)[0][0] < queryVal {
			cur := heap.Pop(h).([]int)
			_, r, c := cur[0], cur[1], cur[2]
			points++
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] {
					heap.Push(h, []int{grid[nr][nc], nr, nc})
					visited[nr][nc] = true
				}
			}
		}
		result[queryIdx] = points
	}
	return result
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
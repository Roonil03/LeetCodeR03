/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	var n int = len(grid)
	return helper(grid, 0, 0, n)
}

func helper(grid [][]int, a, b, n int) *Node {
	if n == 0 {
		return nil
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count += grid[a+i][b+j]
		}
	}
	if count == n*n {
		return &Node{true, true, nil, nil, nil, nil}
	}
	if count == 0 {
		return &Node{false, true, nil, nil, nil, nil}
	}
	n >>= 1
	return &Node{
		Val:         true,
		IsLeaf:      false,
		TopLeft:     helper(grid, a, b, n),
		TopRight:    helper(grid, a, b+n, n),
		BottomLeft:  helper(grid, a+n, b, n),
		BottomRight: helper(grid, a+n, b+n, n),
	}
}
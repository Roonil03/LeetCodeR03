/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func recoverFromPreorder(traversal string) *TreeNode {
//     if len(traversal) == 0 {
// 		return nil
// 	}
// 	var pos int
// 	return helper(traversal, &pos, 0)
// }

// func helper(traversal string, pos *int, depth int) *TreeNode {
// 	start := *pos
// 	for *pos < len(traversal) && traversal[*pos] == '-' {
// 		*pos++
// 	}
// 	if *pos-start != depth {
// 		*pos = start
// 		return nil
// 	}
// 	start = *pos
// 	for *pos < len(traversal) && traversal[*pos] != '-' {
// 		*pos++
// 	}
// 	val, _ := strconv.Atoi(traversal[start:*pos])
// 	node := &TreeNode{Val: val}
// 	node.Left = helper(traversal, pos, depth+1)
// 	node.Right = helper(traversal, pos, depth+1)
// 	return node
// }
func recoverFromPreorder(traversal string) *TreeNode {
	prv := make(map[int]*TreeNode)
	_, v := next(&traversal)
	prv[0] = &TreeNode{v, nil, nil}
	for len(traversal) > 0 {
		l, v := next(&traversal)
		prv[l] = &TreeNode{v, nil, nil}
		if p := prv[l-1]; (*p).Left == nil {
			(*p).Left = prv[l]
		} else {
			(*p).Right = prv[l]
		}
	}
	return prv[0]
}

func next(t *string) (int, int) {
	p, lvl, val := 0, 0, 0
	for (*t)[p] == '-' {
		p, lvl = p+1, lvl+1
	}
	for p < len(*t) && (*t)[p] != '-' {
		p, val = p+1, val*10+int((*t)[p]-'0')
	}
	*t = (*t)[p:]
	return lvl, val
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// if root == nil{
	//     return nil
	// }
	// var helper func(*TreeNode) int
	// var d map[*TreeNode]int
	// d = make(map[*TreeNode]int)
	// helper = func (root *TreeNode)int{
	//     if root == nil{
	//         return 0
	//     }
	//     if _, ex := d[root]; !ex{
	//         d[root] = 1 + max(helper(root.Left), helper(root.Right))
	//     }
	//     return d[root]
	// }
	// l := helper(root.Left)
	// r := helper(root.Right)
	// if l == r{
	//     return root
	// } else if(l < r){
	//     return lcaDeepestLeaves(root.Left)
	// } else{
	//     return lcaDeepestLeaves(root.Right)
	// }
	_, res := dfs(root)
	return res
}

func dfs(root *TreeNode) (int, *TreeNode) {
	left := 0
	lNode := &TreeNode{}
	right := 0
	rNode := &TreeNode{}
	if root.Left != nil {
		left, lNode = dfs(root.Left)
	}
	if root.Right != nil {
		right, rNode = dfs(root.Right)
	}
	if left == right {
		return left + 1, root
	}
	if left > right {
		return left + 1, lNode
	}
	return right + 1, rNode
}
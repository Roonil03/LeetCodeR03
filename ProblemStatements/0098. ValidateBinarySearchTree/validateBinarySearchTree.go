/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
//     if root.Left == nil && root.Right == nil{
//         return true
//     }
//     if (root.Left == nil && root.Right != nil) || (root.Right == nil && root.Left != nil){
//         return false
//     }
//     if root.Left.Val >= root.Val{
//         return false
//     }
//     if root.Right.Val <= root.Val{
//         return false
//     }
//     return isValidBST(root.Left) && isValidBST(root.Right)
// }

func isValidBST(root *TreeNode) bool {
	return v(root, (*TreeNode)(nil), (*TreeNode)(nil))
}

func v(root, m1, m2 *TreeNode) bool {
	if root == nil {
		return true
	}
	if (m1 != nil && root.Val <= m1.Val) || (m2 != nil && root.Val >= m2.Val) {
		return false
	}
	return v(root.Left, m1, root) && v(root.Right, root, m2)
}
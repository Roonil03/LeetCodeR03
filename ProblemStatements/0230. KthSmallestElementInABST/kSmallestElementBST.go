/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	// res := 0
	// count := 0
	// var inorder func(root *TreeNode)
	// inorder = func(root *TreeNode){
	//     if root == nil{
	//         return
	//     }
	//     inorder(root.Left)
	//     count++
	//     if(count == k){
	//         res = root.Val
	//         return
	//     }
	//     inorder(root.Right)
	// }
	// inorder(root)
	// return res

	if root == nil {
		return 0
	}
	leftCount := helper(root.Left)
	if k <= leftCount {
		return kthSmallest(root.Left, k)
	} else if k == leftCount+1 {
		return root.Val
	} else {
		return kthSmallest(root.Right, k-leftCount-1)
	}
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + helper(node.Left) + helper(node.Right)
}
// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// func longestZigZag(root *TreeNode) int {
//     res := 0
//     var dfs func(*TreeNode)(l, r int)
//     dfs = func(node *TreeNode)(l, r int){
//         if node == nil{
//             return -1, -1
//         }
//         _, lr := dfs(node.Left)
//         rl, _ := dfs(node.Right)
//         l = lr + 1
//         r = rl + 1
//         if l > r{
//             res = l
//         }
//         if r > res{
//             res = r
//         }
//         return l, r
//     }
//     dfs(root)
//     return res
// }

func longestZigZag(root *TreeNode) int {
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(node *TreeNode, isLeft bool, count int) int {
	if node == nil {
		return count
	}
	if isLeft {
		return max(dfs(node.Left, true, 0), dfs(node.Right, false, count+1))
	}
	return max(dfs(node.Left, true, count+1), dfs(node.Right, false, 0))
}
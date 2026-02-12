/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    vals := []int{}
    var inorder func(*TreeNode)
    inorder = func(n *TreeNode){
        if n == nil{
            return
        }
        inorder(n.Left)
        vals = append(vals, n.Val)
        inorder(n.Right)
    }
    inorder(root)
    var build func(int, int) *TreeNode
    build = func(a, b int) *TreeNode{
        if a > b{
            return nil
        }
        m := (a + b) / 2
        return &TreeNode{
            Val: vals[m],
            Left: build(a, m-1),
            Right: build(m+1, b),
        }
    }
    return build(0, len(vals) - 1)
}
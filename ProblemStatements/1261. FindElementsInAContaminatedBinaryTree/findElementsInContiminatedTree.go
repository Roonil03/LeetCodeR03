/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	m map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	m := make(map[int]struct{})
	var t func(root *TreeNode, val int)
	t = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		m[val] = struct{}{}
		t(root.Left, 2*val+1)
		t(root.Right, 2*val+2)
	}
	t(root, 0)
	return FindElements{
		m: m,
	}
}

func (this *FindElements) Find(target int) bool {
	if _, ok := this.m[target]; ok {
		return true
	}
	return false
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
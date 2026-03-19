/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, strconv.Itoa(n.Val))
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return strings.Join(res, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	v := strings.Split(data, " ")
	val := make([]int, len(v))
	for i, s := range v {
		val[i], _ = strconv.Atoi(s)
	}
	id := 0
	var build func(int, int) *TreeNode
	build = func(l, h int) *TreeNode {
		if id >= len(val) {
			return nil
		}
		a := val[id]
		if a < l || a > h {
			return nil
		}
		id++
		n := &TreeNode{Val: a}
		n.Left = build(l, a)
		n.Right = build(a, h)
		return n
	}
	return build(-1<<31, 1<<31-1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
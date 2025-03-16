/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	b := root
	for b != nil {
		var head *Node
		var a *Node
		for b != nil {
			if b.Left != nil {
				if a != nil {
					a.Next = b.Left
				} else {
					head = b.Left
				}
				a = b.Left
			}
			if b.Right != nil {
				if a != nil {
					a.Next = b.Right
				} else {
					head = b.Right
				}
				a = b.Right
			}
			b = b.Next
		}
		b = head
	}
	return root
}
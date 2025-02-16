/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var p *ListNode
	s := head
	f := head
	for f != nil && f.Next != nil {
		p = s
		s = s.Next
		f = f.Next.Next
	}
	//if p != nil {
	p.Next = nil
	//}
	n := &TreeNode{Val: s.Val}
	n.Left = sortedListToBST(head)
	n.Right = sortedListToBST(s.Next)
	return n
}
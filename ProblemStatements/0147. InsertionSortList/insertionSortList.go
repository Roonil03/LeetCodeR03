/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	temp := &ListNode{Val: 0, Next: nil}
	cur := head
	for cur != nil {
		next := cur.Next
		prev := temp
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return temp.Next
}
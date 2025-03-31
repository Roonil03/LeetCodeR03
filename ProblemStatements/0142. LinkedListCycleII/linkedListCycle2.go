/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			break
		}
	}
	if f == nil || f.Next == nil {
		return nil
	}
	for head != s {
		head = head.Next
		s = s.Next
	}
	return head
}
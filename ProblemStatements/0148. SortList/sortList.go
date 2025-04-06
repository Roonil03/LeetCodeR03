/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p *ListNode
	s, f := head, head
	for f != nil && f.Next != nil {
		p = s
		s = s.Next
		f = f.Next.Next
	}
	p.Next = nil
	l1 := sortList(head)
	l2 := sortList(s)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	l := &ListNode{Val: 0, Next: nil}
	p := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return l.Next
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := midNode(head)
	nH := mid.Next
	mid.Next = nil
	nH = reverseLinked(nH)
	c1 := head
	c2 := nH
	var f1, f2 *ListNode
	for c1 != nil && c2 != nil {
		f1 = c1.Next
		f2 = c2.Next
		c1.Next = c2
		c2.Next = f1
		c1 = f1
		c2 = f2
	}
}

func midNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseLinked(head *ListNode) *ListNode {
	var prev, cur, forw *ListNode = nil, head, nil
	for cur != nil {
		forw = cur.Next
		cur.Next = prev
		prev = cur
		cur = forw
	}
	return prev
}
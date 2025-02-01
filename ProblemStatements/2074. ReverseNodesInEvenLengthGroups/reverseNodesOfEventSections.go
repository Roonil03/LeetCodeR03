/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	groupSize := 1
	for prev.Next != nil {
		cur := prev.Next
		length := 0
		for i := 0; i < groupSize && cur != nil; i++ {
			length++
			cur = cur.Next
		}
		if length%2 == 0 {
			start := prev.Next
			next := cur
			reversed := reverse(prev.Next, length)
			prev.Next = reversed
			start.Next = next
			prev = start
		} else {
			for i := 0; i < length; i++ {
				prev = prev.Next
			}
		}
		groupSize++
	}
	return dummy.Next
}

func reverse(head *ListNode, length int) *ListNode {
	var prev *ListNode
	cur := head
	for i := 0; i < length; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
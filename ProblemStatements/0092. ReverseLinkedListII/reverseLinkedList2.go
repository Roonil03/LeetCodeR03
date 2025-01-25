/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	temp := &ListNode{0, head}
	prev := temp
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		nn := cur.Next
		cur.Next = nn.Next
		nn.Next = prev.Next
		prev.Next = nn
	}
	return temp.Next
}
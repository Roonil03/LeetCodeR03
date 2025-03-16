/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	vis := make(map[*ListNode]bool)
	c := head
	for c != nil {
		if vis[c] {
			return true
		}
		vis[c] = true
		c = c.Next
	}
	return false
	// if head == nil || head.Next == nil {
	//   return false
	// }
	// s := head
	// f := head.Next
	// for s != f {
	//   if f.Next == nil || f.Next.Next == nil {
	//     return false
	//   }
	//   s = s.Next
	//   f = f.Next.Next
	// }
	// return true
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	// nums := []int{}
	// c := head
	// for c != nil {
	//     nums = append(nums, c.Val)
	//     c = c.Next
	// }
	// res := 0
	// for i := 0; i < len(nums)/2; i++ {
	//     t := nums[i] + nums[len(nums)-1-i]
	//     if t > res {
	//         res = t
	//     }
	// }
	// return res
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		temp := slow.Next
		fast = fast.Next.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}
	res := 0
	for slow != nil {
		sum := slow.Val + prev.Val
		if res < sum {
			res = sum
		}
		slow = slow.Next
		prev = prev.Next
	}
	return res
}
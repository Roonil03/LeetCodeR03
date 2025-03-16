/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var cache = map[*Node]*Node{}

func copyRandomList(head *Node) *Node {
	//     if head == nil{
	//         return nil
	//     }
	//     cur := head
	//     for cur != nil{
	//         n := &Node{Val: cur.Val, Next: cur.Next}
	//         cur.Next = n
	//         cur = n.Next
	//     }
	//     cur = head
	//     for cur !=nil{
	//         if cur.Random != nil{
	//             cur.Next.Random = cur.Random.Next
	//         }
	//         cur = cur.Next.Next
	//     }
	//     cur = head
	//     res := head.Next
	//     for cur != nil{
	//         temp := cur.Next
	//         cur.Next = temp.Next
	//         if temp.Next != nil{
	//             temp.Next = temp.Next.Next
	//         }
	//         cur = cur.Next
	//     }
	//     return res
	// }
	if head == nil {
		return nil
	}
	if node, has := cache[head]; has {
		return node
	}
	node := &Node{Val: head.Val}
	cache[head] = node
	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)
	return node
}
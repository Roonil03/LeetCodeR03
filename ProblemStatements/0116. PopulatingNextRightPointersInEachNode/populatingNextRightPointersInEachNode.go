/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// 	if root == nil{
	//         return nil
	//     }
	//     a := root
	//     for a.Left != nil{
	//         temp := a
	//         for temp != nil{
	//             temp.Left.Next = temp.Right
	//             if temp.Next != nil{
	//                 temp.Right.Next = temp.Next.Left
	//             }
	//             temp = temp.Next
	//         }
	//         a = a.Left
	//     }
	//     return root
	// }
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) != 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if i == n-1 {
				node.Next = nil
			} else {
				node.Next = q[0]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
// func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
//     n := len(status)
//     res := 0
//     haveBox := make([]bool, n)
//     canOpen := make([]bool, n)
//     queue := make([]int, 0)
//     for _, box := range initialBoxes {
//         haveBox[box] = true
//         if status[box] == 1 {
//             canOpen[box] = true
//             queue = append(queue, box)
//         }
//     }
//     for len(queue) > 0 {
//         currentBox := queue[0]
//         queue = queue[1:]
//         res += candies[currentBox]
//         for _, key := range keys[currentBox] {
//             if !canOpen[key] && haveBox[key] {
//                 canOpen[key] = true
//                 queue = append(queue, key)
//             }
//             canOpen[key] = true
//         }
//         for _, newBox := range containedBoxes[currentBox] {
//             haveBox[newBox] = true
//             if !canOpen[newBox] && (status[newBox] == 1 || canOpen[newBox]) {
//                 canOpen[newBox] = true
//                 queue = append(queue, newBox)
//             }
//         }
//     }
//     return res
// }

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	canOpen := make([]bool, n)
	for i := 0; i < n; i++ {
		canOpen[i] = (status[i] == 1)
	}
	hasBox := make([]bool, n)
	used := make([]bool, n)
	q := []int{}
	ans := 0
	for _, box := range initialBoxes {
		hasBox[box] = true
		if canOpen[box] {
			q = append(q, box)
			used[box] = true
			ans += candies[box]
		}
	}
	for len(q) > 0 {
		bigBox := q[0]
		q = q[1:]
		for _, key := range keys[bigBox] {
			canOpen[key] = true
			if !used[key] && hasBox[key] {
				q = append(q, key)
				used[key] = true
				ans += candies[key]
			}
		}
		for _, innerBox := range containedBoxes[bigBox] {
			hasBox[innerBox] = true
			if !used[innerBox] && canOpen[innerBox] {
				q = append(q, innerBox)
				used[innerBox] = true
				ans += candies[innerBox]
			}
		}
	}
	return ans
}
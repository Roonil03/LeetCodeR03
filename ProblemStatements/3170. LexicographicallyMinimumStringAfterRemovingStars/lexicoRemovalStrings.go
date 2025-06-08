// func clearStars(s string) string {
//     chars := []byte(s)
//     var res []byte
//     for i := 0; i < len(chars); i++ {
//         if chars[i] != '*' {
//             res = append(res, chars[i])
//         } else {
//             if len(res) > 0{
//                 y := 0
//                 for j := 0; j < len(res); j++ {
//                     if res[j] < res[y] {
//                         y = j
//                     }
//                 }
//                 res = append(res[:y], res[y+1:]...)
//             }
//         }
//     }
//     return string(res)
// }

// type IntHeap struct {
// 	nums   []int
// 	lessFn func(int, int) bool
// }

// func (h IntHeap) Len() int           { return len(h.nums) }
// func (h IntHeap) Less(i, j int) bool { return h.lessFn(h.nums[i], h.nums[j]) }
// func (h IntHeap) Swap(i, j int)      { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }

// func (h *IntHeap) Push(x any) {
// 	h.nums = append(h.nums, x.(int))
// }

// func (h *IntHeap) Pop() any {
// 	n := len(h.nums)
// 	x := h.nums[n-1]
// 	h.nums = h.nums[0 : n-1]
// 	return x
// }

// func clearStars(s string) string {
// 	hp := IntHeap{
// 		nums: make([]int, 0),
// 		lessFn: func(i1, i2 int) bool {
// 			if s[i1] == s[i2] {
// 				return i1 > i2
// 			}
// 			return s[i1] < s[i2]
// 		},
// 	}
// 	deletedIdx := make(map[int]bool)
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != '*' {
// 			heap.Push(&hp, i)
// 		} else {
// 			idx := heap.Pop(&hp).(int)
// 			deletedIdx[idx] = true
// 		}
// 	}
// 	out := make([]byte, 0, len(s))
// 	for i := 0; i < len(s); i++ {
// 		if !deletedIdx[i] && s[i] != '*' {
// 			out = append(out, s[i])
// 		}
// 	}
// 	return string(out)
// }

func clearStars(s string) string {
	t := []byte(s)
	n := len(s)
	pos := [26][]int{}
	for i := 0; i < n; i++ {
		if t[i] != '*' {
			c := t[i] - 'a'
			pos[c] = append(pos[c], i)
			continue
		}
		var c int
		for c = 0; c < 26; c++ {
			if len(pos[c]) > 0 {
				break
			}
		}
		p := pos[c][len(pos[c])-1]
		pos[c] = pos[c][:len(pos[c])-1]
		t[p] = '*'
	}
	tt := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if t[i] == '*' {
			continue
		}
		tt = append(tt, t[i])
	}
	return string(tt)
}
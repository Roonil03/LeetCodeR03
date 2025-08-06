// func numOfUnplacedFruits(fruits []int, baskets []int) int {
//     // n := len(fruits)
//     // u := make([]bool, n)
//     // un := 0
//     // for i := 0; i < n; i++{
//     //     p := false
//     //     for j := 0; j < n; j++{
//     //         if !u[j] && baskets[j] >= fruits[i]{
//     //             u[j] = true
//     //             p = true
//     //             break
//     //         }
//     //     }
//     //     if !p{
//     //         un++
//     //     }
//     // }
//     // return un
//     n := len(fruits)
// 	type pair struct{ value, index int }
// 	bPairs := make([]pair, n)
// 	for i, v := range baskets {
// 		bPairs[i] = pair{v, i}
// 	}
// 	sort.Slice(bPairs, func(i, j int) bool {
// 		return bPairs[i].value < bPairs[j].value
// 	})
// 	bCap := make([]int, n)
// 	bPos := make([]int, n)
// 	for i := range bPairs {
// 		bCap[i] = bPairs[i].value
// 		bPos[i] = bPairs[i].index
// 	}
// 	st := NewSegmentTree(n)
// 	unplaced := 0
// 	for _, q := range fruits {
//         i := sort.Search(n, func(i int) bool { return bCap[i] >= q })
//         if i == n {
//             unplaced++
//             continue
//         }
//         res := -1
//         for l := i; l < n; l++ {
//             idx := bPos[l]
//             if st.t[idx+st.n] == 1 {
//                 res = l
//                 break
//             }
//         }
//         if res != -1 {
//             st.Update(bPos[res])
//         } else {
//             unplaced++
//         }
//     }
// 	return unplaced
// }

// type SegmentTree struct {
// 	n  int
// 	t  []int
// }

// func NewSegmentTree(n int) *SegmentTree {
// 	t := make([]int, 2*n)
// 	for i := 0; i < n; i++ {
// 		t[n+i] = 1
// 	}
// 	for i := n - 1; i > 0; i-- {
// 		t[i] = t[i<<1] + t[i<<1|1]
// 	}
// 	return &SegmentTree{n, t}
// }

// func (st *SegmentTree) Update(pos int) {
// 	pos += st.n
// 	st.t[pos] = 0
// 	for pos > 1 {
// 		pos >>= 1
// 		st.t[pos] = st.t[pos<<1] + st.t[pos<<1|1]
// 	}
// }

// func (st *SegmentTree) Query(l, r int) int {
// 	l += st.n
// 	r += st.n
// 	for l <= r {
// 		if l&1 == 1 && st.t[l] == 1 {
// 			return l - st.n
// 		}
// 		if r&1 == 0 && st.t[r] == 1 {
// 			return r - st.n
// 		}
// 		l = (l + 1) >> 1
// 		r = (r - 1) >> 1
// 	}
// 	return -1
// }

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(baskets)
	N := 1
	for N <= n {
		N <<= 1
	}
	segTree := make([]int, N<<1)
	for i := 0; i < n; i++ {
		segTree[N+i] = baskets[i]
	}
	for i := N - 1; i > 0; i-- {
		segTree[i] = max(segTree[2*i], segTree[2*i+1])
	}
	count := 0
	for _, fruit := range fruits {
		index := 1
		if segTree[index] < fruit {
			count++
			continue
		}
		for index < N {
			if segTree[2*index] >= fruit {
				index = 2 * index
			} else {
				index = 2*index + 1
			}
		}
		segTree[index] = -1
		for index > 1 {
			index >>= 1
			segTree[index] = max(segTree[2*index], segTree[2*index+1])
		}
	}
	return count
}
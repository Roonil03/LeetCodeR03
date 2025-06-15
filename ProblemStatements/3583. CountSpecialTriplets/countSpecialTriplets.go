const mod = int(1e9 + 7)

type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1),
		n:    n,
	}
}

func (bit *BIT) update(idx, val int) {
	for ; idx <= bit.n; idx += idx & (-idx) {
		bit.tree[idx] = (bit.tree[idx] + val) % mod
	}
}

func (bit *BIT) query(idx int) int {
	sum := 0
	for ; idx > 0; idx -= idx & (-idx) {
		sum = (sum + bit.tree[idx]) % mod
	}
	return sum
}

func specialTriplets(nums []int) int {
	res := 0
	// m := make(map[int]int)
	n := len(nums)
	// for j := 1; j < n-1; j++{
	//     a, b, tar := 0, 0, nums[j] * 2
	//     for i := 0; i < j; i++{
	//         if nums[i] == tar{
	//             a++
	//         }
	//     }
	//     if a == 0{
	//         continue
	//     }
	//     for k := j+1; k < n; k++{
	//         if nums[k] == tar{
	//             b++
	//         }
	//     }
	//     if b > 0{
	//         res = (res + (a * b) % mod) % mod
	//     }
	// // }
	// bit := newBIT(n)
	// s := make([]int, n)
	// copy(s,  nums)
	// sort.Ints(s)
	// comp := make(map[int]int)
	// for i, v := range s{
	//     comp[v] = i + 1
	// }
	// for i := 1; j < n-1; i++{
	//     tar := nums[j] * 2
	//     c, ex := comp[tar]
	//     if !ex{
	//         bit.update(comp[j], 1)
	//         continue
	//     }
	//     a := bit.query(c) - bit.query(c - 1)
	//     b := 0
	//     for k :=
	// }
	vals := make(map[int]bool)
	for _, i := range nums {
		vals[i] = true
		vals[i*2] = true
	}
	s := make([]int, 0, len(vals))
	for v := range vals {
		s = append(s, v)
	}
	sort.Ints(s)
	comp := make(map[int]int)
	for i, v := range s {
		comp[v] = i + 1
	}
	rb := NewBIT(len(s))
	for i := 2; i < n; i++ {
		rb.update(comp[nums[i]], 1)
	}
	lb := NewBIT(len(s))
	lb.update(comp[nums[0]], 1)
	for j := 1; j < n-1; j++ {
		tar := nums[j] * 2
		c, ex := comp[tar]
		if ex {
			a := lb.query(c) - lb.query(c-1)
			b := rb.query(c) - rb.query(c-1)
			res = (res + (a*b)%mod) % mod
		}
		lb.update(comp[nums[j]], 1)
		if j < n-1 {
			rb.update(comp[nums[j+1]], -1)
		}
	}
	return res
}
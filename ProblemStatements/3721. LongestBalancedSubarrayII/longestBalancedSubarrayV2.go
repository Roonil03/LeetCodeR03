func longestBalanced(nums []int) int {
	n := len(nums)
	st := NewSegTree(n)
	id := make([]int, 100001)
	for i := range id {
		id[i] = -1
	}
	res := 0
	for j := 0; j < n; j++ {
		val := nums[j]
		prev := id[val]
		st.Activate(1, 0, n-1, j)
		del := 1
		if val%2 != 0 {
			del = -1
		}
		st.Update(1, 0, n-1, prev+1, j, del)
		l := st.FindLeftmost(1, 0, n-1, 0)
		if l != -1 {
			cur := j - l + 1
			if cur > res {
				res = cur
			}
		}
		id[val] = j
	}
	return res
}

type Node struct {
	minVal int
	maxVal int
	lazy   int
}

type SegTree struct {
	tree []Node
	n    int
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{
		tree: make([]Node, 4*n),
		n:    n,
	}
	for i := range st.tree {
		st.tree[i].minVal = math.MaxInt
		st.tree[i].maxVal = math.MinInt
	}
	return st
}

func (st *SegTree) pushUp(node int) {
	st.tree[node].minVal = min(st.tree[2*node].minVal, st.tree[2*node+1].minVal)
	st.tree[node].maxVal = max(st.tree[2*node].maxVal, st.tree[2*node+1].maxVal)
}

func (st *SegTree) apply(node, val int) {
	if st.tree[node].minVal != math.MaxInt {
		st.tree[node].minVal += val
		st.tree[node].maxVal += val
		st.tree[node].lazy += val
	}
}

func (st *SegTree) pushDown(node int) {
	if st.tree[node].lazy != 0 {
		st.apply(2*node, st.tree[node].lazy)
		st.apply(2*node+1, st.tree[node].lazy)
		st.tree[node].lazy = 0
	}
}

func (st *SegTree) Activate(node, start, end, id int) {
	if start == end {
		st.tree[node].minVal = 0
		st.tree[node].maxVal = 0
		st.tree[node].lazy = 0
		return
	}
	st.pushDown(node)
	mid := (start + end) / 2
	if id <= mid {
		st.Activate(2*node, start, mid, id)
	} else {
		st.Activate(2*node+1, mid+1, end, id)
	}
	st.pushUp(node)
}

func (st *SegTree) Update(node, start, end, l, r, val int) {
	if l > end || r < start {
		return
	}
	if l <= start && end <= r {
		st.apply(node, val)
		return
	}
	st.pushDown(node)
	mid := (start + end) / 2
	st.Update(2*node, start, mid, l, r, val)
	st.Update(2*node+1, mid+1, end, l, r, val)
	st.pushUp(node)
}

func (st *SegTree) FindLeftmost(node, start, end, target int) int {
	if st.tree[node].minVal > target || st.tree[node].maxVal < target {
		return -1
	}
	if start == end {
		return start
	}
	st.pushDown(node)
	mid := (start + end) / 2
	res := st.FindLeftmost(2*node, start, mid, target)
	if res != -1 {
		return res
	}
	return st.FindLeftmost(2*node+1, mid+1, end, target)
}
var trie []Node
var roots []int

const sz = 30

func maxXor(nums []int, k int) int {
	n := len(nums)
	trie = []Node{{lastId: -1}}
	trie[0].lastId = -1
	roots = make([]int, n+1)
	xor := make([]int, n+1)
	roots[0] = insert(0, 0, 0)
	for i := range nums {
		xor[i+1] = xor[i] ^ nums[i]
		roots[i+1] = insert(roots[i], xor[i+1], i+1)
	}
	minQ, maxQ := []int{}, []int{}
	l, res := 0, 0
	for r := range nums {
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[r] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[r] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			l++
			if maxQ[0] < l {
				maxQ = maxQ[1:]
			}
			if minQ[0] < l {
				minQ = minQ[1:]
			}
		}
		cur := query(roots[r], xor[r+1], l)
		res = max(res, cur)
	}
	return res
}

type Node struct {
	children [2]int
	lastId   int
}

func insert(prev, val, id int) int {
	root := len(trie)
	trie = append(trie, trie[prev])
	cur := root
	for i := sz; i >= 0; i-- {
		bit := (val >> i) & 1
		j := len(trie)
		trie = append(trie, trie[trie[cur].children[bit]])
		trie[cur].children[bit] = j
		cur = j
		trie[cur].lastId = id
	}
	return root
}

func query(root, val, minId int) int {
	cur := root
	res := 0
	for i := sz; i >= 0; i-- {
		bit := (val >> i) & 1
		tar := 1 - bit
		cId := trie[cur].children[tar]
		if cId != 0 && trie[cId].lastId >= minId {
			res |= (1 << i)
			cur = cId
		} else {
			cur = trie[cur].children[bit]
		}
	}
	return res
}
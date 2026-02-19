func findMaximumXOR(nums []int) int {
	t := NewTrie()
	t.insert(nums[0])
	res := 0
	for i := 1; i < len(nums); i++ {
		a := t.h1(nums[i])
		res = max(res, a)
		t.insert(nums[i])
	}
	return res
}

type TrieNode struct {
	children [2]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

func (t *Trie) insert(num int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
	}
}

func (t *Trie) h1(num int) int {
	node := t.root
	res := 0
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		opp := 1 - bit
		if node.children[opp] != nil {
			res |= (1 << i)
			node = node.children[opp]
		} else if node.children[bit] != nil {
			node = node.children[bit]
		} else {
			return 0
		}
	}
	return res
}
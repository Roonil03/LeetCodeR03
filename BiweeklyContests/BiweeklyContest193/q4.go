func minInversions(words []string, target []int) int {
	tr := NewTrie[int]()
	pos := make([]int, len(target))
	for i, v := range target {
		pos[v] = i
	}
	for i, v := range words {
		tr.Insert(v, pos[i])
	}
	_, res := dfs(tr.root)
	return int(res)
}

func dfs(node *Node[int]) ([]int, int64) {
	if node == nil {
		return nil, 0
	}
	var cc [][]int
	var temp int64
	for i := range 26 {
		if node.next[i] != nil {
			s, inv := dfs(node.next[i])
			cc = append(cc, s)
			temp += inv
		}
	}
	k := len(cc)
	if k > 0 {
		c := make([][]int64, k)
		for i := range k {
			c[i] = make([]int64, k)
			for j := range k {
				if i != j {
					c[i][j] = Inv(cc[i], cc[j])
				}
			}
		}
		dp := make([]int64, 1<<k)
		for i := 1; i < (1 << k); i++ {
			dp[i] = math.MaxInt64
		}
		for mm := 1; mm < (1 << k); mm++ {
			for i := range k {
				if (mm & (1 << i)) != 0 {
					prev := mm ^ (1 << i)
					cost := int64(0)
					for j := range k {
						if (prev & (1 << j)) != 0 {
							cost += c[j][i]
						}
					}
					dp[mm] = min(dp[prev]+cost, dp[mm])
				}
			}
		}
		temp += dp[(1<<k)-1]
	}
	var mer []int
	for _, v := range cc {
		mer = append(mer, v...)
	}
	if node.term {
		for _, v := range mer {
			if v < node.value {
				temp++
			}
		}
		mer = append(mer, node.value)
	}
	slices.Sort(mer)
	return mer, temp
}

type Node[T any] struct {
	next  [26]*Node[T]
	value T
	term  bool
}

type Trie[T any] struct {
	root *Node[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{root: &Node[T]{}}
}

func (t *Trie[T]) Insert(word string, val T) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			curr.next[c] = &Node[T]{}
		}
		curr = curr.next[c]
	}
	curr.value = val
	curr.term = true
}

func (t *Trie[T]) Search(word string) (T, bool) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			var zero T
			return zero, false
		}
		curr = curr.next[c]
	}
	return curr.value, curr.term
}

func /*(t *Trie[T])*/ Inv(a, b []int) int64 {
	inv := int64(0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else {
			inv += int64(len(a) - i)
			j++
		}
	}
	return inv
}

// O(nlogn . d + n.d.k^2.2^k)
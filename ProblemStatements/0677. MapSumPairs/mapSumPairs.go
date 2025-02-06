type MapSum struct {
	keys map[rune]*MapSum
	sum  int
}

func Constructor() MapSum {
	return MapSum{
		keys: make(map[rune]*MapSum),
		sum:  0,
	}
}

func (this *MapSum) Insert(key string, val int) {
	pointer := this
	for _, char := range key {
		if pointer.keys[char] == nil {
			newKey := Constructor()
			pointer.keys[char] = &newKey
		}
		pointer = pointer.keys[char]
	}
	pointer.sum = val
}

func DFS(node *MapSum) int {
	res := node.sum
	for _, node := range node.keys {
		res += DFS(node)
	}
	return res
}

func (this *MapSum) Sum(prefix string) int {
	for _, char := range prefix {
		if this.keys[char] == nil {
			return 0
		}
		this = this.keys[char]
	}
	return DFS(this)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
type LRUCache struct {
	cap        int
	cache      map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  &Node{},
		tail:  &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.removeNode(node)
		this.helper(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.val = value
		this.removeNode(node)
		this.helper(node)
		return
	}
	newNode := &Node{key: key, val: value}
	this.cache[key] = newNode
	this.helper(newNode)
	if len(this.cache) > this.cap {
		lru := this.tail.prev
		this.removeNode(lru)
		delete(this.cache, lru.key)
	}
}

func (this *LRUCache) helper(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
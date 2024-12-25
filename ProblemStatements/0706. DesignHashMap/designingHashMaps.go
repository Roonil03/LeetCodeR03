type data struct {
	key, val int
}

type MyHashMap struct {
	b    [][]data
	size int
}

func Constructor() MyHashMap {
	return MyHashMap{b: make([][]data, 1000), size: 1000}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	id := this.hash(key)
	for i, p := range this.b[id] {
		if p.key == key {
			this.b[id][i].val = value
			return
		}
	}
	this.b[id] = append(this.b[id], data{key, value})
}

func (this *MyHashMap) Get(key int) int {
	id := this.hash(key)
	for _, p := range this.b[id] {
		if p.key == key {
			return p.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	id := this.hash(key)
	for i, p := range this.b[id] {
		if p.key == key {
			this.b[id] = append(this.b[id][:i], this.b[id][i+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
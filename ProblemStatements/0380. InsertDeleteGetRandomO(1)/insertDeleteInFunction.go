type RandomizedSet struct {
	arr  []int
	size int
	set  map[int]int
}

func Constructor() RandomizedSet {
	arr := make([]int, 0)
	size := 0
	set := make(map[int]int)
	return RandomizedSet{
		arr:  arr,
		size: size,
		set:  set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.set[val] = this.size
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.set[val]
	if !ok {
		return false
	}
	l := this.arr[this.size-1]
	this.arr[id] = l
	this.set[l] = id
	this.arr = this.arr[:this.size-1]
	delete(this.set, val)
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	id := rand.Intn(this.size)
	return this.arr[id]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
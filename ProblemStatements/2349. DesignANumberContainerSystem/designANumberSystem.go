type NumberContainers struct {
	Container map[int]int
	ReIndexes map[int]*treemap.Map
}

func Constructor() NumberContainers {
	return NumberContainers{
		Container: make(map[int]int),
		ReIndexes: make(map[int]*treemap.Map),
	}
}

func (nc *NumberContainers) Change(index int, newNumber int) {
	if oldNumber, exist := nc.Container[index]; exist {
		nc.ReIndexes[oldNumber].Remove(index)
	}
	if nc.ReIndexes[newNumber] == nil {
		nc.ReIndexes[newNumber] = treemap.NewWithIntComparator()
	}
	nc.Container[index] = newNumber
	nc.ReIndexes[newNumber].Put(index, struct{}{})
}

func (nc *NumberContainers) Find(number int) int {
	if nc.ReIndexes[number] == nil || nc.ReIndexes[number].Empty() {
		return -1
	}
	k, _ := nc.ReIndexes[number].Min()
	return k.(int)
}
type MinStack struct {
	s   []int
	m   []int
	top int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.m = append(this.m, val)
	} else {
		this.m = append(this.m, min(val, (this.m)[this.top]))
	}
	this.s = append(this.s, val)
	this.top++
}

func (this *MinStack) Pop() {
	this.s = this.s[:this.top]
	this.m = this.m[:this.top]
	this.top--
	return
}

func (this *MinStack) Top() int {
	return (this.s)[this.top]
}

func (this *MinStack) GetMin() int {
	return (this.m)[this.top]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
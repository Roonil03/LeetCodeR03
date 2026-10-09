type RangeModule struct {
	it [][2]int
}

func Constructor() RangeModule {
	return RangeModule{}
}

func (this *RangeModule) AddRange(left int, right int) {
	i := sort.Search(len(this.it), func(k int) bool {
		return this.it[k][1] >= left
	})
	j := sort.Search(len(this.it), func(k int) bool {
		return this.it[k][0] > right
	})
	if i < j {
		left = min(left, this.it[i][0])
		right = max(right, this.it[j-1][1])
	}
	this.it = slices.Replace(this.it, i, j, [2]int{left, right})
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	i := sort.Search(len(this.it), func(k int) bool {
		return this.it[k][0] > left
	}) - 1
	return i >= 0 && right <= this.it[i][1]
}

func (this *RangeModule) RemoveRange(left int, right int) {
	i := sort.Search(len(this.it), func(k int) bool {
		return this.it[k][1] > left
	})
	j := sort.Search(len(this.it), func(k int) bool {
		return this.it[k][0] >= right
	})
	var buf [2][2]int
	mid := buf[:0]
	if i < j {
		if this.it[i][0] < left {
			mid = append(mid, [2]int{this.it[i][0], left})
		}
		if this.it[j-1][1] > right {
			mid = append(mid, [2]int{right, this.it[j-1][1]})
		}
	}
	this.it = slices.Replace(this.it, i, j, mid...)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
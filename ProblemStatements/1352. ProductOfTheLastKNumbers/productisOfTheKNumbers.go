type ProductOfNumbers struct {
	res []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		res: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.res = []int{1}
	} else {
		this.res = append(this.res, this.res[len(this.res)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.res) {
		return 0
	}
	return this.res[len(this.res)-1] / this.res[len(this.res)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
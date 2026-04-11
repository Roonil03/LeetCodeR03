func minOperations(nums []int) int {
	a := 2
	for _, v := range nums {
		if v > a {
			a = v
		}
	}
	lim := 2*a + 5
	pr := make([]bool, lim+1)
	for i := 2; i <= lim; i++ {
		pr[i] = true
	}
	for i := 2; i*i <= lim; i++ {
		if pr[i] {
			for j := i * i; j <= lim; j += i {
				pr[j] = false
			}
		}
	}
	np := make([]int, lim+1)
	b := -1
	for i := lim; i >= 0; i-- {
		if pr[i] {
			b = i
		}
		np[i] = b
	}
	res := 0
	for i, v := range nums {
		if i%2 == 0 {
			if v <= 2 {
				res += 2 - v
			} else {
				res += np[v] - v
			}
		} else {
			if v <= 1 {
				continue
			}
			if !pr[v] {
				continue
			}
			if v == 2 {
				res += 2
			} else {
				res += 1
			}
		}
	}
	return res
}
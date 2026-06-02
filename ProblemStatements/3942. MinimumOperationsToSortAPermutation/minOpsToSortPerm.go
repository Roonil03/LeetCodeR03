func minOperations(nums []int) int {
	n, p := len(nums), 0
	for i, v := range nums {
		if v == 0 {
			p = i
		}
	}
	a, b := true, true
	for i := range n {
		if nums[(i+1)%n] != (nums[i]+1)%n {
			a = false
		}
		if nums[(i+1)%n] != (nums[i]-1+n)%n {
			b = false
		}
	}
	res := -1
	if a {
		res = min(p, n-p+2)
	}
	if b {
		if c := min(p+2, n-p); res == -1 || c < res {
			res = c
		}
	}
	return res
}
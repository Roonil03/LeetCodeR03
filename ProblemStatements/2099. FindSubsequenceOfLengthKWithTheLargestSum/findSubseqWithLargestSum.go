func maxSubsequence(nums []int, k int) []int {
	type p struct {
		v, i int
	}
	a := make([]p, len(nums))
	for i, v := range nums {
		a[i] = p{v, i}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].v == a[j].v {
			return a[i].i < a[j].i
		}
		return a[i].v > a[j].v
	})
	a = a[:k]
	sort.Slice(a, func(i, j int) bool { return a[i].i < a[j].i })
	res := make([]int, k)
	for i := range a {
		res[i] = a[i].v
	}
	return res
}
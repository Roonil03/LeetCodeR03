func maxSum(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n)
	copy(s, nums)
	sort.Ints(s)
	res := int64(nums[0])
	for i := range n {
		a, b := make([]int, 0, n), make([]int, n)
		copy(b, s)
		sum := int64(0)
		for j := i; j < n; j++ {
			v := nums[j]
			sum += int64(v)
			q := sort.SearchInts(a, v)
			a = append(a, 0)
			copy(a[q+1:], a[q:])
			a[q] = v
			q = sort.SearchInts(b, v)
			copy(b[q:], b[q+1:])
			b = b[:len(b)-1]
			g := int64(0)
			for c := 0; c < k && c < len(a) && c < len(b); c++ {
				if b[len(b)-1-c] > a[c] {
					g += int64(b[len(b)-1-c] - a[c])
				} else {
					break
				}
			}
			res = max(res, sum+g)
		}
	}
	return res
}
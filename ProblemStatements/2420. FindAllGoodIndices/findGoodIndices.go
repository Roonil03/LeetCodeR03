func goodIndices(nums []int, k int) []int {
	res := []int{}
	l, r := 0, 0
outer:
	for i := k; i < len(nums)-k; i++ {
		j := max(l, i-k)
		for j < i-1 {
			if nums[j] < nums[j+1] {
				i = j + k
				continue outer
			}
			j++
		}
		j = max(r, i+1)
		for j < k+i {
			if nums[j] > nums[j+1] {
				i = j - 1
				continue outer
			}
			j++
		}
		l = i - 1
		r = i + k - 1
		res = append(res, i)
	}
	return res
}
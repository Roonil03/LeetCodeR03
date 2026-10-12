func resilientSubarray(nums []int, k int) int {
	res, n := -1, len(nums)
	for i := range nums {
		r := (nums[i]%k + k) % k
		j := i + 1
		for j < n && (nums[j]%k+k)%k == r {
			j++
		}
		if l := j - i; l > res {
			d := 1
			if r != 0 {
				a, b := k, r
				for b != 0 {
					a, b = b, a%b
				}
				d = k / a
			}
			res = max(res, 1+(l-1)/d*d)
		}
		i = j
	}
	return res
}
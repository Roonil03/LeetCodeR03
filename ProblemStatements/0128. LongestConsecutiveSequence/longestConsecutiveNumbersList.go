func longestConsecutive(nums []int) int {
	a := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		a[nums[i]] = struct{}{}
	}
	res := 0
	for n := range a {
		if _, ok := a[n-1]; !ok {
			v := 1
			for {
				if _, ok = a[n+v]; ok {
					v++
					continue
				}
				res = max(res, v)
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
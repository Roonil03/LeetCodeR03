func longestCommonPrefix(arr1 []int, arr2 []int) int {
	pre := make(map[int]struct{})
	for _, v := range arr1 {
		for v > 0 {
			pre[v] = struct{}{}
			v /= 10
		}
	}
	res := 0
	for _, v := range arr2 {
		l := 0
		for i := v; i > 0; i /= 10 {
			l++
		}
		if l <= res {
			continue
		}
		for v > 0 {
			if _, ex := pre[v]; ex {
				if l > res {
					res = l
				}
				break
			}
			v /= 10
			l--
		}
	}
	return res
}
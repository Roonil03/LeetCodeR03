func minimumDistance(nums []int) int {
	mp := make(map[int][]int)
	res := (1 << 32) - 1
	for i, v := range nums {
		mp[v] = append(mp[v], i)
		if j, ok := mp[v]; ok && len(j) > 2 {
			for k := 0; k <= len(j)-3; k++ {
				dist := abs(j[k]-j[k+1]) + abs(j[k+1]-j[k+2]) + abs(j[k+2]-j[k])
				res = min(res, dist)
			}
		}
	}
	if res == (1<<32)-1 {
		return -1
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
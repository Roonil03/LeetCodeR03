func minJumps(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	g := make(map[int][]int)
	for i, v := range arr {
		g[v] = append(g[v], i)
	}
	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	res := 0
	for len(q) > 0 {
		sz := len(q)
		for i := range sz {
			cur := q[i]
			if cur == n-1 {
				return res
			}
			for _, v := range g[arr[cur]] {
				if !vis[v] {
					vis[v] = true
					q = append(q, v)
				}
			}
			delete(g, arr[cur])
			if cur+1 < n && !vis[cur+1] {
				vis[cur+1] = true
				q = append(q, cur+1)
			}
			if cur-1 >= 0 && !vis[cur-1] {
				vis[cur-1] = true
				q = append(q, cur-1)
			}
		}
		q = q[sz:]
		res++
	}
	return 0
}
func maximumMEX(nums []int) []int {
	n := len(nums)
	a := make([]int, n+2)
	for _, v := range nums {
		if v <= n {
			a[v]++
		}
	}
	m := 0
	for a[m] > 0 {
		m++
	}
	var res []int
	vis := make([]bool, n+2)
	for i := 0; i < n; {
		if m == 0 {
			res = append(res, 0)
			i++
			continue
		}
		t, c := m, 0
		for ; i < n; i++ {
			v := nums[i]
			if v < t && !vis[v] {
				vis[v] = true
				c++
			}
			if v <= n {
				a[v]--
				if a[v] == 0 && v < m {
					m = v
				}
			}
			if c == t {
				i++
				break
			}
		}
		for j := range t {
			vis[j] = false
		}
		res = append(res, t)
	}
	return res
}
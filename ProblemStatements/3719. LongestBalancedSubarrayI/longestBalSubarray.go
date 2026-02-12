func longestBalanced(nums []int) int {
    n := len(nums)
	res := 0
	for l := 0; l < n; l++ {
		vis := make(map[int]bool, n-l)
		e, o := 0, 0
		for r := l; r < n; r++ {
			x := nums[r]
			if !vis[x] {
				vis[x] = true
				if x%2 == 0 {
					e++
				} else {
					o++
				}
			}
			if e == o {
				res = max(res, r - l +1)
			}
		}
	}
	return res
}

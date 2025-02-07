func sumOfPowers(nums []int, k int) int {
	const mod = 1000000007
	n := len(nums)
	sort.Ints(nums)
	pwrs := make(map[int]bool)
	dp := make([][][]int, 50)
	for i := range dp {
		dp[i] = make([][]int, 50)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var dfs func(i, p, k, pwr int, eq bool, n []int) int
	dfs = func(i, p, k, pwr int, eq bool, n []int) int {
		if k == 0 {
			if eq {
				return pwr
			}
			return 0
		}
		if i >= len(n) {
			return 0
		}
		eqInt := 0
		if eq {
			eqInt = 1
		}
		if dp[p][k][eqInt] == -1 {
			dp[p][k][eqInt] = dfs(i+1, p, k, pwr, eq, n)
			if abs(n[i]-n[p]) >= pwr {
				newEq := eq || (n[i]-n[p] == pwr)
				val := dfs(i+1, i, k-1, pwr, newEq, n)
				dp[p][k][eqInt] = (dp[p][k][eqInt] + val) % mod
			}
		}
		return dp[p][k][eqInt]
	}
	var res int64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] != nums[j] {
				pwr := nums[j] - nums[i]
				if !pwrs[pwr] {
					pwrs[pwr] = true
					for x := range dp {
						for y := range dp[x] {
							for z := range dp[x][y] {
								dp[x][y][z] = -1
							}
						}
					}
					val := dfs(0, n-1, k, pwr, false, nums)
					res = (res + int64(val)) % int64(mod)
				}
			}
		}
	}
	return int(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
package main

func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		}
		res := 1
		for j := i + 1; j <= i+d && j < n && arr[i] > arr[j]; j++ {
			res = max(res, 1+dfs(j))
		}
		for j := i - 1; j >= i-d && j >= 0 && arr[i] > arr[j]; j-- {
			res = max(res, 1+dfs(j))
		}
		dp[i] = res
		return res
	}
	r := 0
	for i := range arr {
		r = max(r, dfs(i))
	}
	return r
}

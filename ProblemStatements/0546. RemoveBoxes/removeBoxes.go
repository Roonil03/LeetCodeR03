func removeBoxes(boxes []int) int {
    n := len(boxes)
    dp := make([][][]int, n)
    for i := range dp{
        dp[i] = make([][]int, n)
        for j := range dp[i]{
            dp[i][j] = make([]int, n)
        }
    }
    var dfs func(l, r, k int) int
    dfs = func(l, r, k int) int{
        if l > r{
            return 0
        }
        if dp[l][r][k] != 0{
            return dp[l][r][k]
        }
        a, c := l, k
        for l < r && boxes[l] == boxes[l + 1]{
            l++
            k++
        }
        res := (k + 1) * (k + 1) + dfs(l + 1, r, 0)
        for i := l + 1; i <= r; i++{
            if boxes[i] == boxes[l]{
                res = max(res, dfs(l + 1, i - 1, 0) + dfs(i, r, k + 1))
            }
        }
        dp[a][r][c] = res
        return res
    }
    return dfs(0, n - 1, 0)
}
func minScoreTriangulation(values []int) int {
    n := len(values)
    dp := make([][]int, n)
    for i := range dp{
        dp[i] = make([]int, n)
    }
    for i := 2; i < n; i++{
        for j := 0; j + i < n; j++{
            k := j + i
            cur := 1 << 30
            for m := j + 1; m < k; m++{
                temp := dp[j][m] + dp[m][k] + values[j] * values[m] * values[k]
                if temp < cur{
                    cur = temp
                }
            }
            dp[j][k] = cur
        }
    }
    return dp[0][n - 1]
}
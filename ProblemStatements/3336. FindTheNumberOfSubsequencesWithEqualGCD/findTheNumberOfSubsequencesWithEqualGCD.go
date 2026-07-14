func subsequencePairCount(nums []int) int {
    m := 0
    mod := int(1e9+7)
    for _, v := range nums{
        m = max(v, m)
    }
    g := make([][]int, m + 1)
    for i := range g{
        g[i] = make([]int, m + 1)
        for j := range g[i]{
            a, b := i, j
            for b != 0{
                a, b = b, a%b
            }
            g[i][j] = a
        }
    }
    s := m + 1
    dp := make([]int, s * s)
    dp[0] = 1
    for _, v := range nums{
        ndp := make([]int, s * s)
        for i := range (m + 1){
            for j := range (m + 1){
                x := dp[i * s + j]
                if x == 0{
                    continue
                }
                ndp[i *s + j] = (ndp[i*s+j] + x) % mod
                ni, nj := v, v
                if i != 0{
                    ni = g[i][v]
                }
                if j != 0{
                    nj = g[j][v]
                }
                ndp[ni * s + j] = (ndp[ni * s + j] + x) % mod
                ndp[i * s + nj] = (ndp[i * s + nj] + x) % mod
            }
        }
        dp = ndp
    }
    res := 0
    for i := 1; i <= m; i++{
        res = (res + dp[i * s + i]) % mod
    }
    return res
}
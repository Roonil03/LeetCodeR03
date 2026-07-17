func gcdValues(nums []int, queries []int64) []int {
    m := slices.Max(nums)
    freq := make([]int64, m + 1)
    for _, v := range nums{
        freq[v]++
    }
    pre := make([]int64, m + 1)
    for i := m; i>= 1; i--{
        temp := int64(0)
        for j := i; j <= m; j+= i{
            temp += freq[j]
        }
        p := temp * (temp -  1) / 2
        for j := 2 * i; j <= m; j += i{
            p -= pre[j]
        }
        pre[i] = p
    }
    p := make([]int64, m+1)
    for i := 1; i <= m;i++{
        p[i] = p[i-1] + pre[i]
    }
    res := make([]int, len(queries))
    for i, v := range queries{
        res[i] = sort.Search(m+1, func(j int) bool{
            return p[j] > v
        })
    }
    return res
}
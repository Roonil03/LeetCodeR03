func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    comp := make([]int, n)
    c := 0
    for i := 1; i < n; i++{
        if nums[i] - nums[i -1] > maxDiff{
            c++
        }
        comp[i] = c
    }
    res := make([]bool, len(queries))
    for i, q := range queries{
        res[i] = comp[q[0]] == comp[q[1]]
    }
    return res
}
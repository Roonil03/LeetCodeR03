func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    i, j := 0, 0
    res := 0
    for i < len(g) && j < len(s){
        if s[j] >= g[i]{
            res++
            i++
            j++
        } else{
            j++
        }
    }
    return res
}
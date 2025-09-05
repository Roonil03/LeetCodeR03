func mergeTriplets(triplets [][]int, target []int) bool {
    x, y, z := target[0], target[1], target[2]
    cur := []int{0, 0, 0}
    for _,t  := range triplets{
        if t[0] <= x && t[1] <= y && t[2] <= z{
            if t[0] > cur[0]{
                cur[0] = t[0]
            }
            if t[1] > cur[1]{
                cur[1] = t[1]
            }
            if t[2] > cur[2]{
                cur[2] = t[2]
            }
        }
    }
    return cur[0] == x && cur[1] == y && cur[2] == z
}
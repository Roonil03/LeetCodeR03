func sequentialDigits(low int, high int) []int {
    var res []int
    for i := 1; i <= 9; i++{
        for j := 1; j <= 10 - i; j++{
            n := 0
            for k := range i{
                n = n * 10 + j + k
            }
            if n >= low && n <= high{
                res = append(res, n)
            }
        }
    }
    return res
}
func sumAndMultiply(n int) int64 {
    x, sum := 0, 0
    for _, v := range strconv.Itoa(n){
        if v > '0' && v <= '9'{
            d := int(v - '0')
            x, sum = x * 10 + d, sum + d
        }
    }
    return int64(x) * int64(sum)
}
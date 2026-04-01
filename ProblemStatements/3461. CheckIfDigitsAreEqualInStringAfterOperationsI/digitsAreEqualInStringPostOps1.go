func hasSameDigits(s string) bool {
    n := len(s) - 2
    d := make([]int, len(s))
    for i := range len(s){
        d[i] = int(s[i] - '0')
    }
    v2, v5, c := make([]int, n+1), make([]int, n+1), make([]int, n+1)
    for i := 1; i <= n; i++{
        if i % 2 == 0{
            v2[i] = v2[i/2] + 1
            v5[i] = v5[i/2]
            c[i] = c[i/2]
        } else if i % 5 == 0{
            v2[i] = v2[i/5]
            v5[i] = v5[i/5] + 1
            c[i] = c[i/5]
        } else{
            c[i] = i % 10
        }
    }
    inv := [10]int{}
    inv[1] = 1
    inv[3] = 7
    inv[7] = 3
    inv[9] = 9
    rem, e2, e5 := 1, 0, 0
    l, r := d[0], d[1]
    for i := range n{
        x, y := n - i, i + 1
        e2 += v2[x] - v2[y]
        e5 += v5[x] - v5[y]
        rem = (rem * c[x]) % 10
        rem = (rem * inv[c[y]]) % 10
        a := h1(rem, e2, e5)
        l = (l + a * d[i+1]) % 10
        r = (r + a * d[i+2]) % 10
        // fmt.Println(l, r)
    }
    return l == r
}

func h1(rem, e2, e5 int) int {
	if e5 > 0 {
		if e2 > 0 {
			return 0
		}
		return (rem * 5) % 10
	}
	return (rem * h2(e2)) % 10
}

func h2(e int) int {
	if e == 0 {
		return 1
	}
	switch e % 4 {
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 8
	default:
		return 6
	}
}
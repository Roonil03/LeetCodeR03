func reverseParentheses(s string) string {
    n := len(s)
    p := make([]int, n)
    st := make([]int, 0, n)
    for i := range n{
        if s[i] == '('{
            st = append(st, i)
        } else if s[i] == ')'{
            j := st[len(st) - 1]
            st = st[:len(st) - 1]
            p[i] = j
            p[j] = i
        }
    }
    res := make([]byte, 0, n)
    for i, d := 0, 1; i>= 0 && i < n; i += d{
        if s[i] == '(' || s[i] == ')'{
            i = p[i]
            d = -d
        } else{
            res = append(res, s[i])
        }
    }
    return string(res)
}
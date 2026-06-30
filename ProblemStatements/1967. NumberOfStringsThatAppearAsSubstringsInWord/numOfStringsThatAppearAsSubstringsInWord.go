func numOfStrings(patterns []string, word string) int {
    type node struct{
        n [26]*node
        fail *node
        id []int
    }
    root := &node{}
    for i, v := range patterns{
        cur := root
        for j := range v{
            ch := v[j] - 'a'
            if cur.n[ch] == nil{
                cur.n[ch] = &node{}
            }
            cur = cur.n[ch]
        }
        cur.id = append(cur.id, i)
    }
    var q []*node
    for i := range 26{
        if root.n[i] != nil{
            root.n[i].fail = root
            q = append(q, root.n[i])
        } else{
            root.n[i] = root
        }
    }
    for len(q) > 0{
        cur := q[0]
        q = q[1:]
        for i := range 26{
            if cur.n[i] != nil{
                cur.n[i].fail = cur.fail.n[i]
                cur.n[i].id = append(cur.n[i].id, cur.fail.n[i].id...)
                q = append(q, cur.n[i])
            } else{
                cur.n[i] = cur.fail.n[i]
            }
        }
    }
    f := make([]bool, len(patterns))
    cur := root
    for i := range word{
        cur = cur.n[word[i] - 'a']
        for _, v := range cur.id{
            f[v] = true
        }
    }
    res := 0
    for _, v := range f{
        if v{
            res++
        }
    }
    return res
}
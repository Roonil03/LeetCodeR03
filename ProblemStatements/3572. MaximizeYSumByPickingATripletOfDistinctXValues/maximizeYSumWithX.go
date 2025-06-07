func maxSumDistinctTriplet(x []int, y []int) int {
	n := len(y)
	// q := make([]p, n)
	// for i:= 0; i < n; i++{
	//     q[i] = p{y: y[i], x : x[i], id : i}
	// }
	// s := merge(q)
	// res := -1
	// for i:= 0; i < n-2; i++{
	//     for j := i+1; j < n-1; j++{
	//         for k := j+1; k < n; k++{
	//             p1,p2,p3 := s[i],s[j],s[k]
	//             if p1.x != p2.x && p2.x != p3.x && p1.x != p3.x{
	//                 res = p1.y + p2.y + p3.y
	//                 return res
	//             }
	//         }
	//     }
	// }
	// return res
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		q := x[i]
		w := y[i]
		if w > m[q] {
			m[q] = w
		}
	}
	var z []int
	for _, v := range m {
		z = append(z, v)
	}
	if len(z) < 3 {
		return -1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(z)))
	res := 0
	for i := 0; i < 3; i++ {
		res += z[i]
	}
	return res
}

type p struct {
	y  int
	x  int
	id int
}

func merge(a []p) []p {
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	l := merge(a[:m])
	r := merge(a[m:])
	return help(l, r)
}

func help(l, r []p) []p {
	res := make([]p, 0, len(l)+len(r))
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i].y >= r[j].y {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, r[j:]...)
	return res
}
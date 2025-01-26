func countMentions(numberOfUsers int, events [][]string) []int {
	m := numberOfUsers
	//n := len(events)
	l := [][3]interface{}{}
	for _, e := range events {
		if e[0] == "MESSAGE" {
			t, _ := strconv.Atoi(e[1])
			l = append(l, [3]interface{}{t, 2, e[2]})
		} else {
			t, _ := strconv.Atoi(e[1])
			l = append(l, [3]interface{}{t, 0, e[2]})
			l = append(l, [3]interface{}{t + 60, 1, e[2]})
		}
	}
	d := make(map[int][][]string)
	for _, v := range l {
		t := v[0].(int)
		x := v[1].(int)
		r := v[2].(string)
		if _, ok := d[t]; !ok {
			d[t] = make([][]string, 3)
		}
		d[t][x] = append(d[t][x], r)
	}
	u := make(map[int]bool)
	for i := 0; i < m; i++ {
		u[i] = true
	}
	a := make([]int, m)
	keys := make([]int, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, t := range keys {
		for _, user := range d[t][1] {
			uid, _ := strconv.Atoi(user)
			u[uid] = true
		}
		for _, user := range d[t][0] {
			uid, _ := strconv.Atoi(user)
			if u[uid] {
				delete(u, uid)
			}
		}
		for _, x := range d[t][2] {
			if x == "ALL" {
				for i := 0; i < m; i++ {
					a[i]++
				}
			} else if x == "HERE" {
				for i := range u {
					a[i]++
				}
			} else {
				s := strings.Fields(x)
				for _, i := range s {
					idx, _ := strconv.Atoi(i[2:])
					a[idx]++
				}
			}
		}
	}
	return a
}
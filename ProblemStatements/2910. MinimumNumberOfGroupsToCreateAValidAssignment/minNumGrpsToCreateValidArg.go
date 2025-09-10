func minGroupsForValidAssignment(balls []int) int {
	freq := map[int]int{}
	for _, x := range balls {
		freq[x]++
	}
	mf := len(balls)
	for _, v := range freq {
		if v < mf {
			mf = v
		}
	}
	cd := func(a, b int) int {
		return (a + b - 1) / b
	}
	res := -1
	for i := mf; i >= 1; i-- {
		tot, ok := 0, true
		for _, v := range freq {
			l := cd(v, i+1)
			r := v / i
			if l > r {
				ok = false
				break
			}
			tot += l
		}
		if ok {
			res = tot
			break
		}
	}
	return res
}
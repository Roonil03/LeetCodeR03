func longestBalanced(s string) int {
	res := 0
	c := []byte{'a', 'b', 'c'}
	for i := 1; i < 8; i++ {
		subs := []byte{}
		for j := 0; j < 3; j++ {
			if (i>>j)&1 == 1 {
				subs = append(subs, c[j])
			}
		}
		res = max(res, h(s, subs))
	}
	return res
}

func h(s string, sub []byte) int {
	t := make([]bool, 256)
	for _, b := range sub {
		t[b] = true
	}
	m := make(map[string]int)
	counts := make([]int, 256)
	m[h1(counts, sub)] = -1
	res := 0
	lastStart := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if !t[ch] {
			for j := lastStart; j < i; j++ {
				counts[s[j]]--
			}
			lastStart = i + 1
			m = make(map[string]int)
			m[h1(counts, sub)] = i
			continue
		}
		counts[ch]++
		ok := true
		for _, b := range sub {
			if counts[b] == 0 {
				ok = false
				break
			}
		}
		k := h1(counts, sub)
		if f, ok2 := m[k]; ok2 {
			if ok {
				if d := i - f; d > res {
					res = d
				}
			}
		} else {
			m[k] = i
		}
	}
	return res
}

func h1(counts []int, subset []byte) string {
	if len(subset) == 1 {
		return "s"
	}
	var sb []byte
	for i := 0; i < len(subset)-1; i++ {
		diff := counts[subset[i]] - counts[subset[i+1]]
		if diff < 0 {
			sb = append(sb, '-')
			diff = -diff
		}
		if diff == 0 {
			sb = append(sb, '0')
		} else {
			var temp []byte
			for diff > 0 {
				temp = append(temp, byte('0'+(diff%10)))
				diff /= 10
			}
			for j := len(temp) - 1; j >= 0; j-- {
				sb = append(sb, temp[j])
			}
		}
		sb = append(sb, ',')
	}
	return string(sb)
}
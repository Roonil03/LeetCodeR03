func minimumGroups(words []string) int {
	mp := make(map[string]struct{})
	for _, v := range words {
		e, o := []byte{}, []byte{}
		for i := range v {
			if i%2 == 0 {
				e = append(e, v[i])
			} else {
				o = append(o, v[i])
			}
		}
		mp[h1(e)+"_"+h1(o)] = struct{}{}
	}
	return len(mp)
}

func h1(s []byte) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	i, j := 0, 1
	for i < n && j < n {
		k := 0
		for k < n && s[(i+k)%n] == s[(j+k)%n] {
			k++
		}
		if k == n {
			break
		}
		if s[(i+k)%n] > s[(j+k)%n] {
			i += k + 1
			if i == j {
				i++
			}
		} else {
			j += k + 1
			if i == j {
				j++
			}
		}
	}
	if j < i {
		i = j
	}
	r := make([]byte, n)
	for k := range r {
		r[k] = s[(i+k)%n]
	}
	return string(r)
}
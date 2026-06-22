func goodIntegers(l int64, r int64, k int) int64 {
	res := h1(strconv.FormatInt(r, 10), k)
	if l > 0 {
		res -= h1(strconv.FormatInt(l-1, 10), k)
	}
	return res
}

func h1(s string, k int) int64 {
	m := make([][11][2]int64, len(s))
	for i := range m {
		for j := range m[i] {
			m[i][j] = [2]int64{-1, -1}
		}
	}
	var dp func(int, int, int, int) int64
	dp = func(i, p, t, lz int) int64 {
		if i == len(s) {
			return 1
		}
		if t == 0 && m[i][p][lz] != -1 {
			return m[i][p][lz]
		}
		lim := 9
		if t == 1 {
			lim = int(s[i] - '0')
		}
		res := int64(0)
		for j := 0; j <= lim; j++ {
			if lz == 0 && (j-p > k || p-j > k) {
				continue
			}
			nt, nlz, np := 0, 0, j
			if t == 1 && j == lim {
				nt = 1
			}
			if lz == 1 && j == 0 {
				nlz, np = 1, 10
			}
			res += dp(i+1, np, nt, nlz)
		}
		if t == 0 {
			m[i][p][lz] = res
		}
		return res
	}
	return dp(0, 10, 1, 1)
}
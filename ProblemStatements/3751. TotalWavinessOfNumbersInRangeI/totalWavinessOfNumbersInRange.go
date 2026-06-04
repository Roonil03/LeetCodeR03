func totalWaviness(num1 int, num2 int) int {
	h1 := func(n int) int {
		if n < 100 {
			return 0
		}
		var dig []int
		for n > 0 {
			dig = append(dig, n%10)
			n /= 10
		}
		slices.Reverse(dig)
		var mem, mems [20][11][11][2]int
		for i := range mem {
			for j := range mem[i] {
				for k := range mem[i][j] {
					mem[i][j][k] = [2]int{-1, -1}
					mems[i][j][k] = [2]int{-1, -1}
				}
			}
		}
		var dp func(id, p1, p2, t, start int) (int, int)
		dp = func(id, p1, p2, t, start int) (int, int) {
			if id == len(dig) {
				return 1, 0
			}
			if t == 0 && mem[id][p1][p2][start] != -1 {
				return mem[id][p1][p2][start], mems[id][p1][p2][start]
			}
			lim := 9
			if t == 1 {
				lim = dig[id]
			}
			count, sum := 0, 0
			for i := 0; i <= lim; i++ {
				a := 0
				if t == 1 && i == lim {
					a = 1
				}
				b := start
				if i > 0 {
					b = 1
				}
				wave := 0
				if start == 1 && p2 != 10 {
					if (p1 > p2 && p1 > i) || (p1 < p2 && p1 < i) {
						wave = 1
					}
				}
				np2 := p2
				if start == 1 {
					np2 = p1
				}
				np1 := 10
				if b == 1 {
					np1 = i
				}
				c, s := dp(id+1, np1, np2, a, b)
				count += c
				sum += s + wave*c
			}
			if t == 0 {
				mem[id][p1][p2][start] = count
				mems[id][p1][p2][start] = sum
			}
			return count, sum
		}
		_, s := dp(0, 10, 10, 1, 0)
		return s
	}
	return h1(num2) - h1(num1-1)
}
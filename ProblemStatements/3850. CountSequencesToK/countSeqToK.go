type f struct {
	n, d int64
}

func countSequences(nums []int, k int64) int {
	dp := make(map[f]int)
	dp[f{1, 1}] = 1
	for _, x := range nums {
		nt := make(map[f]int)
		for i, v := range dp {
			nt[i] += v
			mul := h1(i.n*int64(x), i.d)
			nt[mul] += v
			if x != 0 {
				div := h1(i.n, i.d*int64(x))
				nt[div] += v
			}
		}
		dp = nt
	}
	tar := f{k, 1}
	return dp[tar]
}

func h1(n, d int64) f {
	if n == 0 {
		return f{0, 1}
	}
	if d < 0 {
		n = -n
		d = -d
	}
	g := gcd(abs(n), d)
	return f{n / g, d / g}
}

func abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}


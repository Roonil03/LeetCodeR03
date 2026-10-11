func maxPrimeSubset(n int, s int) []int {
	if n < 2 || s < 2 {
		return []int{}
	}
	pp := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		pp[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if pp[i] {
			for j := i * i; j <= n; j += i {
				pp[j] = false
			}
		}
	}
	var pr []int
	for i := n; i >= 2; i-- {
		if pp[i] {
			pr = append(pr, i)
		}
	}
	dp := make([][]int, s+1)
	dp[0] = []int{}
	for _, v := range pr {
		for j := s; j >= v; j-- {
			if dp[j-v] != nil {
				if dp[j] == nil || len(dp[j-v])+1 <= len(dp[j]) {
					sub := make([]int, len(dp[j-v]), len(dp[j-v])+1)
					copy(sub, dp[j-v])
					dp[j] = append(sub, v)
				}
			}
		}
	}
	for i := s; i > 0; i-- {
		if dp[i] != nil {
			res := dp[i]
			for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
				res[l], res[r] = res[r], res[l]
			}
			return res
		}
	}
	return []int{}
}

// Question solved post the contest time by myself.

/* Greedy forward algorithm is used with dynamic programming. Each state dp[i][w] is denoting min number of primes needed to achieve
that sum using a subset of suffix of primes. Seive of Eratosthenes is used to calculate the primes before hand.
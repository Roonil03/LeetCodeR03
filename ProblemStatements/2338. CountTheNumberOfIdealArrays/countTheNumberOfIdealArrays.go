//CODE THAT DOES NOT WORK:
// const MOD int = 1e9 + 7

// var comb [10001][14]int
// var cnt [10001][14]int

// func idealArrays(n int, maxValue int) int {
//     // comb := make([][]int, n + 1)
//     // for i := range comb{
//     //     comb[i] = make([]int, n + 1)
//     //     comb[i][0] = 1
//     //     for j := 1; j <= i; j++{
//     //         comb[i][j] =  (comb[i-1][j-1] + comb[i-1][j]) % MOD
//     //     }
//     // }
//     // dp := make([]int, maxValue + 1)
//     // for i := 1 ; i <= maxValue; i++{
//     //     dp[i] = 1
//     // }
//     // res := 0
//     // for l := 1; l <= n; l++{
//     //     next := make([]int, maxValue + 1)
//     //     for i := 1; i <= maxValue; i++{
//     //         res = (res + dp[i] * comb[n-1][l - 1]) % MOD
//     //         for k := 2 * i; k <= maxValue; k+= i{
//     //             next[k] = (next[k] + dp[i]) % MOD
//     //         }
//     //     }
//     //     dp = next
//     // }
//     // return res
// 	if comb[1][1] == 0 {
// 		precomp()
// 	}
// 	res := 0
// 	for i := 1; i <= maxValue; i++ {
// 		for bars := 0; bars < min(14, n) && cnt[i][bars] != 0; bars++ {
// 			res = (res + (cnt[i][bars] * comb[n-1][bars])%MOD) % MOD
// 		}
// 	}
// 	return res
// }

// func precomp() {
// 	for s := 1; s <= 10000; s++ {
// 		for r := 0; r < 14; r++ {
// 			if r == 0 {
// 				comb[s][r] = 1
// 			} else {
// 				comb[s][r] = (comb[s-1][r-1] + comb[s-1][r]) % MOD
// 			}
// 		}
// 	}
// 	for div := 1; div <= 10000; div++ {
// 		cnt[div][0]++
// 		for i := 2 * div; i <= 10000; i += div {
// 			for bars := 0; cnt[div][bars] != 0; bars++ {
// 				cnt[i][bars+1] = (cnt[i][bars+1] + cnt[div][bars]) % MOD
// 			}
// 		}
// 	}
// }

//CODE THAT FINISHES WITH MASSIVE TIME DELAY:
// const mod = 1e9+7

// func idealArrays(n int, maxValue int) int {
//     dp := make([][]int, 15)
//     for i := range dp {
//         dp[i] = make([]int, maxValue+1)
//     }
//     m := make(map[int][]int)
//     for i := 1; i <= maxValue; i++ {
//         j := i * 2
//         for j <= maxValue {
//             m[j] = append(m[j], i)
//             j += i
//         }
//     }
//     for i := 1; i <= maxValue; i++ {
//         dp[1][i] = 1
//     }
//     for i := 2; i <= n && i <= 14; i++ {
//         for j := 1; j <= maxValue; j++ {
//             for _, k := range m[j] {
//                 dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
//             }
//         }
//     }
//     for i := 1; i <= n && i <= 14; i++ {
//         for j := 1; j <= maxValue; j++ {
//             dp[i][0] = (dp[i][0] + dp[i][j]) % mod
//         }
//     }
//     memo := make([][]int, n+1)
//     for i := range memo {
//         memo[i] = make([]int, 15)
//     }
//     res := 0
//     for i := 1; i <= n && i <= 14; i++ {
//         res += nCk(n-1, i-1, memo) * dp[i][0]%mod
//         res %= mod
//     }
//     return res
// }

// func nCk(n, k int, memo [][]int) int {
//     if k == 0 {
//         return 1
//     }
//     if n == 0 {
//         return 0
//     }
//     if memo[n][k] != 0 {
//         return memo[n][k]
//     }
//     memo[n][k] = nCk(n-1, k, memo) + nCk(n-1, k-1, memo)
//     memo[n][k] %= mod
//     return memo[n][k]
// }

//CODE THAT HOPEFULLY COMPLETES WITH 0%:
const (
	MOD   int = 1e9 + 7
	MAX_N     = 10010
	MAX_P     = 15
)

var (
	c     [MAX_N + MAX_P][MAX_P + 1]int
	sieve [MAX_N]int
	ps    [MAX_N][]int
)

func initialize() {
	if c[0][0] != 0 {
		return
	}
	for i := 2; i < MAX_N; i++ {
		if sieve[i] == 0 {
			for j := i; j < MAX_N; j += i {
				if sieve[j] == 0 {
					sieve[j] = i
				}
			}
		}
	}
	for i := 2; i < MAX_N; i++ {
		x := i
		for x > 1 {
			p := sieve[x]
			cnt := 0
			for x%p == 0 {
				x /= p
				cnt++
			}
			ps[i] = append(ps[i], cnt)
		}
	}
	c[0][0] = 1
	for i := 1; i < MAX_N+MAX_P; i++ {
		c[i][0] = 1
		for j := 1; j <= MAX_P && j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
		}
	}
}

func idealArrays(n int, maxValue int) int {
	initialize()
	ans := 0
	for x := 1; x <= maxValue; x++ {
		mul := 1
		for _, p := range ps[x] {
			mul = mul * c[n+p-1][p] % MOD
		}
		ans = (ans + mul) % MOD
	}
	return ans
}
const MOD = 1e9 + 7

// func numOfSubarrays(arr []int) int {
//     n := len(arr)
//     p := make([]int, n+1)
//     o, e := 0, 1
//     res := 0
//     for i := 0; i < n; i++{
//         p[i+1] = (p[i] + arr[i]) % MOD
//         if p[i+1] %2 == 0{
//             res = (res + o)%MOD
//             e++
//         } else {
//             res = (res + e)%MOD
//             o++
//         }
//     }
//     return res
// }

func numOfSubarrays(arr []int) int {
	result := 0
	prefixParity, oddPrefixCount, evenPrefixCount := 0, 0, 1
	for _, num := range arr {
		prefixParity = (prefixParity + num) % 2
		if prefixParity == 0 {
			result += oddPrefixCount
			evenPrefixCount++
		} else {
			result += evenPrefixCount
			oddPrefixCount++
		}
		result %= MOD
	}
	return result
}
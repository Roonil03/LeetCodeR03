func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	m := make(map[int64]int64)
	m[0] = 1
	var res, sum int64 = 0, 0
	for _, i := range nums {
		if i%modulo == k {
			sum++
		}
		sum %= int64(modulo)
		// t := (sum - int64(k) + int64(modulo)) % int64(modulo)
		// res += m[t]
		// m[sum]++
		if c, exists := m[(sum-int64(k)+int64(modulo))%int64(modulo)]; exists {
			res += c
		}
		m[sum]++
	}
	return res
}
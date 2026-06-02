func countValidSubsets(parent []int, nums []int, k int) int {
	n := len(parent)
	dp0 := make([]int, n*k)
	dp1 := make([]int, n*k)
	for i := range n {
		dp0[i*k] = 1
		dp1[i*k+((nums[i]%k)+k)%k] = 1
	}
	const mod = int(1e9 + 7)
	ndp0 := make([]int, k)
	ndp1 := make([]int, k)
	for i := n - 1; i > 0; i-- {
		p := parent[i]
		p0 := dp0[p*k : p*k+k]
		p1 := dp1[p*k : p*k+k]
		i0 := dp0[i*k : i*k+k]
		i1 := dp1[i*k : i*k+k]
		for j := 0; j < k; j++ {
			ndp0[j] = 0
			ndp1[j] = 0
		}
		for r1 := range k {
			v0 := p0[r1]
			v1 := p1[r1]
			if v0 == 0 && v1 == 0 {
				continue
			}
			for r2 := range k {
				s := i0[r2] + i1[r2]
				if s >= mod {
					s -= mod
				}
				nr := r1 + r2
				if nr >= k {
					nr -= k
				}
				if v0 > 0 && s > 0 {
					ndp0[nr] = int((int64(ndp0[nr]) + int64(v0)*int64(s)) % int64(mod))
				}
				if v1 > 0 && i0[r2] > 0 {
					ndp1[nr] = int((int64(ndp1[nr]) + int64(v1)*int64(i0[r2])) % int64(mod))
				}
			}
		}
		for j := 0; j < k; j++ {
			p0[j] = ndp0[j]
			p1[j] = ndp1[j]
		}
	}
	return (dp0[0] + dp1[0] - 1 + mod) % mod
}

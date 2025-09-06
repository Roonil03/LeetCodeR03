func minOperations(queries [][]int) int64 {
	//     a := 0
	//     for _, q := range queries{
	//         if q[1] > a {
	//             a = q[1]
	//         }
	//     }
	//     ops := make([]int64, a+1)
	//     for i := 1; i <= a; i++{
	//         count := 0
	//         x := i
	//         for x > 0{
	//             x >>= 2
	//             count++
	//         }
	//         ops[i] = int64(count)
	//     }
	//     pre := make([]int64, a+1)
	//     for i := 1; i <= a; i++{
	//         pre[i] = pre[i-1] + ops[i]
	//     }
	//     var tot int64 = 0
	//     for _, q := range queries{
	//         l, r := q[0], q[1]
	//         rs := pre[r] - pre[l-1]
	//         tot += (rs + 1) / 2
	//     }
	//     return tot
	// }

	var ans int64 = 0
	for _, q := range queries {
		l, r := q[0], q[1]
		var S int64 = 0
		dMax := 0
		for k := 1; k <= 31; k++ {
			low := int64(1) << (k - 1)
			high := (int64(1) << k) - 1
			if low > int64(r) {
				break
			}
			a := max(int64(l), low)
			b := min(int64(r), high)
			if a > b {
				continue
			}
			cnt := b - a + 1
			stepsForK := (k + 1) / 2
			S += cnt * int64(stepsForK)
			if stepsForK > dMax {
				dMax = stepsForK
			}
		}
		ops := max(int64(dMax), (S+1)/2)
		ans += ops
	}
	return ans
}
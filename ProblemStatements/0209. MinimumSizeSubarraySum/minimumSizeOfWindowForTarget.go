// func minSubArrayLen(s int, nums []int) int {
//     n := len(nums)
//     res := int(^uint(0) >> 1)
//     p := make([]int, n+1)
//     for i := 1; i <= n; i++ {
//         p[i] = p[i-1] + nums[i-1]
//     }
//     for i := n; i >= 0 && p[i] >= s; i-- {
//         j := sort.SearchInts(p, p[i]-s+1)
//         if c := i - j + 1; c < res {
//             res = c
//         }
//     }
//     if res == int(^uint(0)>>1) {
//         return 0
//     }
//     return res
// }

func minSubArrayLen(s int, a []int) int {
	if a == nil || len(a) == 0 {
		return 0
	}
	i, j, sum, m := 0, 0, 0, math.MaxInt
	for j < len(a) {
		sum += a[j]
		j++
		for sum >= s {
			m = min(m, j-i)
			sum -= a[i]
			i++
		}
	}
	if m == math.MaxInt {
		return 0
	}
	return m
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
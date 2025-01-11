func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	f := make(map[int]int)
	for j := 0; j < k; j++ {
		f[nums[j]]++
	}
	ans[0] = calXSum(f, x)
	for i := 1; i <= n-k; i++ {
		f[nums[i-1]]--
		if f[nums[i-1]] == 0 {
			delete(f, nums[i-1])
		}
		f[nums[i+k-1]]++
		ans[i] = calXSum(f, x)
	}
	return ans
}

func calXSum(f map[int]int, x int) int {
	p := make([][2]int, 0, len(f))
	for v, c := range f {
		p = append(p, [2]int{c, v})
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] > p[j][1]
		}
		return p[i][0] > p[j][0]
	})
	s := 0
	for i := 0; i < x && i < len(p); i++ {
		s += p[i][1] * p[i][0]
	}
	return s
}

// func findXSum(nums []int, k int, x int) []int {
//     n := len(nums)
//     ans := make([]int, n - k + 1)
//     for i := 0;  i <= n-k; i++{
//         freq:=make(map[int]int)
//         for j := i; j < i+k; j++{
//             freq[nums[j]]++
//         }
//         pairs := make([][2]int, 0)
//         for num, count := range freq{
//             pairs = append(pairs, [2]int{num, count})
//         }
//         sort.Slice(pairs, func(i, j int) bool {
//                 if pairs[i][1] == pairs[j][1] {
//                     return pairs[i][0] > pairs[j][0]
//             }
//             return pairs[i][1] > pairs[j][1]
//         })
//         sum := 0
//         dis := len(pairs)
//         ele := min(x, dis)
//         for i:= 0; i < ele ; i++{
//             sum += pairs[i][0] * pairs[i][1]
//         }
//         ans[i] = sum
//     }
//     return ans
// }
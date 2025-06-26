func findKDistantIndices(nums []int, key int, k int) []int {
	// n := len(nums)
	// m := make([]bool, n)
	// for i, v := range nums{
	//     if v == key{
	//         s := int(math.Max(0, float64(i - k)))
	//         e := int(math.Min(float64(n - 1), float64(i+k)))
	//         for j := s; j <= e; j++{
	//             m[j] = true
	//         }
	//     }
	// }
	// res := []int{}
	// for i, v := range m{
	//     if v{
	//         res = append(res, i)
	//     }
	// }
	res := make([]int, 0)
	start, end := 0, 0
	for {
		if nums[end] == key {
			for ; start <= end+k && start < len(nums); start++ {
				res = append(res, start)
				if nums[start] == key {
					end = start
				}
			}
			end = start
		} else {
			end++
		}
		if end >= len(nums) {
			break
		}
		if end-start > k {
			start++
		}
	}
	return res
}
func longestNiceSubarray(nums []int) int {
	// a := nums[0]
	// res := 1
	// l := 0
	// for r := 1; r < len(nums); r++{
	//     for l != r && nums[r] & a != 0{
	//         a -= nums[l]
	//         l++
	//     }
	//     res = max(res, r - l + 1)
	//     a += nums[r]
	// }
	// return res
	a, b, res := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		for (a & nums[i]) != 0 {
			a ^= nums[b]
			b++
		}
		a |= nums[i]
		res = max(res, i-b+1)
	}
	return res
}
// func maxAbsoluteSum(nums []int) int {
//     k := func(arr []int)int{
//         a, b := 0,0
//         for _, num := range arr{
//             a = int(math.Max(float64(num), float64(a) + float64(num)))
//             b = int(math.Max(float64(b), float64(a)))
//         }
//         return b
//     }
//     res := k(nums)
//     neg := make([]int, len(nums))
//     for i, num:= range nums{
//         neg[i] = -num
//     }
//     min := k(neg)
//     if res > min{
//         return res
//     }
//     return -min
// }

// func maxAbsoluteSum(nums []int) int {
//     curPositiveSum, curNegativeSum := 0, 0
//     maxAbsSum := 0
//     for _, num := range nums {
//         curPositiveSum = max(0, curPositiveSum + num)
//         curNegativeSum = max(0, curNegativeSum - num)
//         maxAbsSum = max(maxAbsSum, max(curPositiveSum, curNegativeSum))
//     }
//     return maxAbsSum
// }

// func maxAbsoluteSum(nums []int) int {
//     maxSum, minSum, currMax, currMin := 0, 0, 0, 0
// 	for _, num := range nums {
// 		currMax = max(num, currMax+num)
// 		maxSum = max(maxSum, currMax)
// 		currMin = min(num, currMin+num)
// 		minSum = min(minSum, currMin)
// 	}
// 	return max(abs(maxSum), abs(minSum))
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func abs(a int) int {
// 	return int(math.Abs(float64(a)))
// }

func maxAbsoluteSum(nums []int) int {
	r := 0
	p := 0
	n := 0
	for _, num := range nums {
		p += num
		n += num
		if p < 0 {
			p = 0
		}
		if n > 0 {
			n = 0
		}
		if p > r {
			r = p
		}
		if -n > r {
			r = -n
		}
	}
	return r
}
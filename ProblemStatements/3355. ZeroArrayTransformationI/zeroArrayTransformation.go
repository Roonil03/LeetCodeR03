// func isZeroArray(nums []int, queries [][]int) bool {
//     n := len(nums)
// 	diff := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		diff[i] = nums[i]
// 		if i > 0 {
// 			diff[i] -= nums[i-1]
// 		}
// 	}
// 	for _, q := range queries {
// 		l, r := q[0], q[1]
// 		diff[l] -= 1
// 		if r+1 < n {
// 			diff[r+1] += 1
// 		}
// 	}
// 	prefix := 0
// 	for i := 0; i < n; i++ {
// 		prefix += diff[i]
// 		if prefix != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	delta := make([]int, n+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		delta[l]++
		if r+1 < n {
			delta[r+1]--
		}
	}
	newDelta := make([]int, n)
	newDelta[0] = delta[0]
	for i := 1; i < n; i++ {
		newDelta[i] = newDelta[i-1] + delta[i]
	}
	for i := 0; i < n; i++ {
		if nums[i]-newDelta[i] > 0 {
			return false
		}
	}
	return true
}
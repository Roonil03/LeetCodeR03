// func pivotArray(nums []int, pivot int) []int {
//     n1 := []int{}
// 	n2 := []int{}
// 	n3 := []int{}
// 	for _, num := range nums {
// 		if num < pivot {
// 			n1 = append(n1, num)
// 		} else if num == pivot {
// 			n2 = append(n2, num)
// 		} else {
// 			n3 = append(n3, num)
// 		}
// 	}
// 	return append(append(n1, n2...), n3...)
// }

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		ans[i] = pivot

		if nums[i] < pivot {
			ans[l] = nums[i]
			l++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[len(nums)-1-i] > pivot {
			ans[r] = nums[len(nums)-1-i]
			r--
		}
	}
	return ans
}
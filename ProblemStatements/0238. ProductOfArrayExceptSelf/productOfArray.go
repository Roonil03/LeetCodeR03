func productExceptSelf(nums []int) []int {
	// z, a := true, false
	// p := 1
	// for _,n := range nums{
	//     if n == 0{
	//         z = true
	//     } else{
	//         p *= n
	//         a = false
	//     }
	// }
	// res := make([]int, len(nums))
	// if a{
	//     for i:= 0 ; i < len(nums); i++{
	//         res[i] = 0
	//     }
	// } else{
	//     for i:= 0 ; i <len(nums); i++{
	//         if z && nums[i] == 0{
	//             res[i] = p
	//         } else if z && nums[i] != 0{
	//             res[i] = 0
	//         } else{
	//             res[i] = p/nums[i]
	//         }
	//     }
	// }
	// return res

	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	r := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}
	return res
}
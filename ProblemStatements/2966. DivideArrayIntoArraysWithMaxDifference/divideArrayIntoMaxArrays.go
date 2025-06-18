func divideArray(nums []int, k int) [][]int {
	//     sort.Ints(nums)
	//     n := len(nums)
	//     mem := make(map[int][][][]int)
	//     var h func(int)([][]int, bool)
	//     h = func(i int) ([][]int, bool){
	//         if i == n{
	//             return [][]int{}, true
	//         }
	//         if sol, ok := mem[i]; ok{
	//             if len(sol) > 0{
	//                 return sol[0], true
	//             }
	//             return nil, false
	//         }
	//         for j := 1 + i; j < n; j++{
	//             for m := j + 1; m < n; m++{
	//                 if nums[k] - nums[i] > k{
	//                     break
	//                 }
	//                 a := []int{nums[i], nums[j], nums[m]}
	//                 if r, p := h(m + 1); p{
	//                     sol := append([][]int{a}, r...)
	//                     mem[i] = [][][]int{sol}
	//                     return sol, true
	//                 }
	//             }
	//         }
	//         mem[i] = [][][]int{}
	//         return nil, false
	//     }
	//     if sol, ok := h(0); ok{
	//         return sol
	//     }
	//     return [][]int{}
	n := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i += 3 {
		if i+2 < n && nums[i+2]-nums[i] <= k {
			res = append(res, []int{nums[i], nums[i+1], nums[i+2]})
		} else {
			return [][]int{}
		}
	}
	return res
}


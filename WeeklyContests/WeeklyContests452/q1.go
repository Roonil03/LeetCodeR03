func checkEqualPartitions(nums []int, target int64) bool {
	// n := len(nums)
	// var tp int64 = 1
	// for _, i:= range nums{
	//     tp *= int64(i)
	// }
	// if tp != target * target{
	//     return false
	// }
	// sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// suf := make([]int, n+1)
	// suf[n] = 1
	// for i := n-1; i >= 0; i--{
	//     suf[i] = suf[i+1] * nums[i]
	// }
	// return bt(nums, 0, 1, target, suf)
	var back func(id int, pro1, pro2 int64, used1, used2 bool) bool
	back = func(id int, pro1, pro2 int64, used1, used2 bool) bool {
		if id == len(nums) {
			return used1 && used2 && pro1 == target && pro2 == target
		}
		a := int64(nums[id])
		if pro1 <= target && (pro1 == 1 || target%pro1 == 0) {
			if back(id+1, pro1*a, pro2, true, used2) {
				return true
			}
		}
		if pro2 <= target && (pro2 == 1 || target%pro2 == 0) {
			if back(id+1, pro1, pro2*a, used1, true) {
				return true
			}
		}
		return false
	}
	return back(0, 1, 1, false, false)
}

// func bt(nums []int, id, pro int, target int64, suf []int) bool{
//     if int64(pro) == target{
//         return true
//     }
//     if id >= len(nums){
//         return false
//     }
//     if int64(pro) * int64(suf[id]) < target{
//         return false
//     }
//     a := pro * nums[id]
//     if int64(a) <= target && bt(nums, id + 1, a, target, suf){
//         return true
//     }
//     return bt(nums, id + 1, pro, target, suf)
// }
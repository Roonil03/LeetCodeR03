func maxFrequencyElements(nums []int) int {
	m := make([]int, 101)
	res, a := 0, 0
	for _, i := range nums {
		m[i]++
	}
	for i := range m {
		if a < m[i] {
			a = m[i]
		}
	}
	for i := range m {
		if a == m[i] {
			res += m[i]
		}
	}
	return res
	// a, b := 0, 1
	// for i := range nums{
	//     m[nums[i]]++
	//     if(m[nums[i]] > a){
	//         a = m[nums[i]]
	//         b = 1
	//     } else if(m[nums[i]] == a){
	//         b++
	//     }
	// }
	// return a * b
}
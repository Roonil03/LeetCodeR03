func repairCars(ranks []int, cars int) int64 {
	//     l, r := int64(0), int64(1e14)
	//     for l < r{
	//         m := l + (r - l) / 2
	//         if helper(ranks, cars, m){
	//             r = m
	//         } else{
	//             l = m+1
	//         }
	//     }
	//     return l
	// }

	// func helper(ranks []int, cars int, m int64)bool{
	//     sum := 0
	//     for _, rank := range ranks{
	//         a := int(math.Floor(math.Sqrt(float64(m)/float64(rank))))
	//         sum += a
	//     }
	//     return sum >= cars
	mx := slices.Max(ranks)
	left, right := int64(0), int64(mx)*int64(cars)*int64(cars)
	check := func(v int64) bool {
		sum := int64(0)
		for _, rank := range ranks {
			sum += int64(math.Sqrt(float64(v / int64(rank))))
			if sum >= int64(cars) {
				return true
			}
		}
		return false
	}
	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
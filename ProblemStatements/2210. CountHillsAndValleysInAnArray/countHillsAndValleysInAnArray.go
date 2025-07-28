func countHillValley(nums []int) int {
	n := len(nums)
	a, b := nums[0], 0
	d := []bool{false, false}
	for i := 1; i < n; i++ {
		for i < n && a == nums[i] {
			i++
		}
		if i == n {
			break
		}
		big := nums[i] > a
		if big {
			d[1] = true
		} else {
			d[0] = true
		}
		if d[0] && d[1] {
			b++
			if big {
				d[0] = false
			} else {
				d[1] = false
			}
		}
		a = nums[i]
	}
	return b
}
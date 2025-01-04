func countValidSelections(nums []int) int {
	n := len(nums)
	count := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if simulate(nums, i, true) {
				count++
			}
			if simulate(nums, i, false) {
				count++
			}
		}
	}

	return count
}

func simulate(nums []int, pos int, right bool) bool {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)

	for pos >= 0 && pos < n {
		if temp[pos] == 0 {
			if right {
				pos++
			} else {
				pos--
			}
		} else {
			temp[pos]--
			right = !right
			if right {
				pos++
			} else {
				pos--
			}
		}
	}

	for _, val := range temp {
		if val != 0 {
			return false
		}
	}

	return true
}
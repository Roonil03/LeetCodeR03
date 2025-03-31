func minimumIndex(nums []int) int {
	major, cnt := nums[0], 0
	for _, n := range nums {
		if n == major {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			major = n
			cnt = 1
		}
	}
	majCount := 0
	for _, n := range nums {
		if n == major {
			majCount++
		}
	}
	cnt = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == major {
			cnt++
		}
		if cnt*2 > (i+1) && (majCount-cnt)*2 > (len(nums)-i-1) {
			return i
		}
	}
	return -1
}
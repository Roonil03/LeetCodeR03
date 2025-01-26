func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	kCount := 0
	for _, x := range nums {
		if x == k {
			kCount++
		}
	}
	answer := kCount
	for x := 1; x <= 50; x++ {
		currK := 0
		myMin := 0
		xCount := 0
		for i := 0; i < n; i++ {
			if nums[i] == x {
				xCount++
			}
			if nums[i] == k {
				currK++
			}
			myMin = min(myMin, xCount-currK)
			answer = max(answer, xCount-myMin+kCount-currK)
		}
	}
	return answer
}
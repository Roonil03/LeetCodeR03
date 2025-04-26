func countSubarrays(nums []int, minK int, maxK int) int64 {
	var count int64 = 0
	start, minPos, maxPos := -1, -1, -1
	for i, num := range nums {
		if num < minK || num > maxK {
			start = i
			minPos = -1
			maxPos = -1
			continue
		}
		if num == minK {
			minPos = i
		}
		if num == maxK {
			maxPos = i
		}
		if minPos != -1 && maxPos != -1 {
			count += int64(max(0, min(minPos, maxPos)-start))
		}
	}
	return count
}
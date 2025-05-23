func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var sum, count, sac int64 = 0, 0, math.MaxInt64
	for _, n := range nums {
		sum += max(int64(n)^int64(k), int64(n))
		if (int64(n) ^ int64(k)) > int64(n) {
			count += 1
		}
		sac = min(sac, abs(int64(n)-(int64(n)^int64(k))))
	}
	if count%2 == 1 {
		return sum - sac
	}
	return sum
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
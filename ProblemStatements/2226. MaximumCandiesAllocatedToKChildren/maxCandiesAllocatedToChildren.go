func maximumCandies(candies []int, k int64) int {
	left, right := 1, 0
	sum := int64(0)
	for _, candy := range candies {
		if candy > right {
			right = candy
		}
		sum += int64(candy)
	}
	if sum < k {
		return 0
	}
	for left < right {
		mid := left + (right-left+1)/2
		count := int64(0)
		for _, candy := range candies {
			count += int64(candy / mid)
		}
		if count >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
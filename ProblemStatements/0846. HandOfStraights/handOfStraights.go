func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, num := range hand {
		count[num]++
	}
	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		c := count[k]
		if c > 0 {
			for i := k; i < k+groupSize; i++ {
				if count[i] < c {
					return false
				}
				count[i] -= c
			}
		}
	}
	return true
}
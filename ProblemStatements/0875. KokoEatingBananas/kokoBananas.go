func minEatingSpeed(piles []int, h int) int {
	a := 0
	for _, pile := range piles {
		if pile > a {
			a = pile
		}
	}
	l, r := 1, a
	for l < r {
		mid := (l + r) / 2
		if eat(mid, h, piles) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func eat(k, h int, piles []int) bool {
	ts := 0
	for _, pile := range piles {
		ts += (pile / k)
		pile = pile % k
		if pile > 0 {
			ts++
		}
	}
	return ts <= h
}
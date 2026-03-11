func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	r := minutesToTest/minutesToDie + 1
	res := 0
	a := 1
	for a < buckets {
		a *= r
		res++
	}
	return res
}
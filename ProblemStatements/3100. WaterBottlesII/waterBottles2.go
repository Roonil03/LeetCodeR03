func maxBottlesDrunk(numBottles int, numExchange int) int {
	tot, emp, k := numBottles, numBottles, numExchange
	for emp >= k {
		emp -= k
		tot++
		emp++
		k++
	}
	return tot
}
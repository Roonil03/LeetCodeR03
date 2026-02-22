func countPrimeSetBits(left int, right int) int {
	test := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true, 23: true, 29: true,
	}
	res := 0
	for i := left; i <= right; i++ {
		b := bits.OnesCount(uint(i))
		if test[b] {
			res++
		}
	}
	return res
}
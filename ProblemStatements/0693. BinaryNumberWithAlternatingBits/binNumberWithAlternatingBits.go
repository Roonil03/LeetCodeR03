func hasAlternatingBits(n int) bool {
	var x uint32 = uint32(n) ^ (uint32(n) >> 1)
	return (x & (x + 1)) == 0
}
func mirrorDistance(n int) int {
	rev := 0
	for i := n; i != 0; i /= 10 {
		rev = 10*rev + i%10
	}
	rev -= n
	if rev >= 0 {
		return rev
	}
	return -rev
}
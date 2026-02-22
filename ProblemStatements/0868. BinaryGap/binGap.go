func binaryGap(n int) int {
	res, lPos, cur := 0, -1, 0
	for n > 0 {
		if n&1 == 1 {
			if lPos != -1 {
				dist := cur - lPos
				res = max(res, dist)
			}
			lPos = cur
		}
		n >>= 1
		cur++
	}
	return res
}
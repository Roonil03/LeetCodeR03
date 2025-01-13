func canThreePartsEqualSum(arr []int) bool {
	t := 0
	for _, num := range arr {
		t += num
	}
	if t%3 != 0 {
		return false
	}
	tar := t / 3
	cs, p := 0, 0
	for _, nums := range arr {
		cs += nums
		if cs == tar {
			p++
			cs = 0
		}
	}
	return p >= 3
}
func minAdjacentSwaps(nums []int, a int, b int) int {
	mod := int(1e9 + 7)
	var x, y, res int
	for _, v := range nums {
		if v < a {
			res = (res + x + y) % mod
		} else if v > b {
			y++
		} else {
			res = (res + y) % mod
			x++
		}
	}
	return res
}
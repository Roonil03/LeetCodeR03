func findGCD(nums []int) int {
	a, b := nums[0], nums[0]
	for _, v := range nums[1:] {
		a = min(v, a)
		b = max(v, b)
	}
	return gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
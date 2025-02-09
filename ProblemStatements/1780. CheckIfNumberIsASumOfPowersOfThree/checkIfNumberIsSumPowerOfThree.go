func checkPowersOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 == 2 {
		return false
	}
	return checkPowersOfThree(n / 3)
}
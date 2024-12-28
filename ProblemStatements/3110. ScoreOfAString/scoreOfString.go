func scoreOfString(s string) int {
	n := len(s) - 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(int(s[i]) - int(s[i+1]))
	}
	return sum
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
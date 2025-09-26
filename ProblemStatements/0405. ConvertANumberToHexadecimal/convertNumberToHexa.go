func toHex(num int) string {
	if num < 0 {
		num += 1 << 32
	}
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	res := ""
	for num > 0 {
		x := num % 16
		res = string(hex[x]) + res
		num /= 16
	}
	return res
}
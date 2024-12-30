package main

func addBinary(a string, b string) string {
	res := ""
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		res = string(carry%2+'0') + res
		carry /= 2
	}
	return res
}

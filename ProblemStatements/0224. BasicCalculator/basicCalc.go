func calculate(s string) int {
	stack := []int{}
	sign := 1
	res := 0
	num := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if unicode.IsDigit(rune(char)) {
			num = 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--
			res += sign * num
		} else if char == '+' {
			sign = 1
		} else if char == '-' {
			sign = -1
		} else if char == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1
		} else if char == ')' {
			res *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

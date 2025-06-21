func calPoints(operations []string) int {
	s := []int{}
	for _, op := range operations {
		switch op {
		case "+":
			n := len(s)
			s = append(s, s[n-1]+s[n-2])
		case "D":
			n := len(s)
			s = append(s, 2*s[n-1])
		case "C":
			s = s[:len(s)-1]
		default:
			num, _ := strconv.Atoi(op)
			s = append(s, num)
		}
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
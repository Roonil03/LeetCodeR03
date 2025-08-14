func largestGoodInteger(num string) string {
	res := ""
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			temp := num[i : i+3]
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
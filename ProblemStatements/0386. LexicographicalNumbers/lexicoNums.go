func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	num := 1
	for len(res) < n {
		res = append(res, num)
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return res
}
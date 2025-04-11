func countSymmetricIntegers(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		if i < 100 {
			if i%10 == i/10 {
				res++
			}
		} else if i < 10000 && i >= 1000 {
			if (i/1000 + (i/100)%10) == ((i / 10 % 10) + i%10) {
				res++
			}
		}
	}
	return res
}
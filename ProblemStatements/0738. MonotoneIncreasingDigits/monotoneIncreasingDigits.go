func monotoneIncreasingDigits(n int) int {
	prev := 10
	mult := 1
	mm := 0
	for temp := n; temp > 0; temp /= 10 {
		d := temp % 10
		if d > prev {
			mm = mult
			prev = d - 1
		} else {
			prev = d
		}
		mult *= 10
	}
	if mm > 0 {
		return (n/mm-1)*mm + mm - 1
	}
	return n
}
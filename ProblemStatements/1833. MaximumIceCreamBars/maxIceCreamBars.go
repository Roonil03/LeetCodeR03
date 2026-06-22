func maxIceCream(costs []int, coins int) int {
	freq := make([]int, 100001)
	for _, v := range costs {
		freq[v]++
	}
	res := 0
	for i := 1; i <= 100000; i++ {
		if freq[i] == 0 {
			continue
		}
		b, a := freq[i], coins/i
		if b > a {
			b = a
		}
		res += b
		coins -= b * i
	}
	return res
}
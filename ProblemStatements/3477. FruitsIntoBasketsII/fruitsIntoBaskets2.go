func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	u := make([]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		p := false
		for j := 0; j < n; j++ {
			if !u[j] && baskets[j] >= fruits[i] {
				u[j] = true
				p = true
				break
			}
		}
		if !p {
			res++
		}
	}
	return res
}
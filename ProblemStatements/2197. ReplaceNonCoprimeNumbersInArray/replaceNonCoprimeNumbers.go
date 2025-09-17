func replaceNonCoprimes(nums []int) []int {
	res := []int{}
	for _, n := range nums {
		res = append(res, n)
		for len(res) > 1 {
			a := res[len(res)-1]
			b := res[len(res)-2]
			g := gcd(a, b)
			if g > 1 {
				res = res[:len(res)-2]
				lcm := a / g * b
				res = append(res, lcm)
			} else {
				break
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
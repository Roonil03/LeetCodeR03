func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func prod(arr []int) *big.Int {
	res := big.NewInt(1)
	for _, num := range arr {
		res.Mul(res, big.NewInt(int64(num)))
	}
	return res
}

func maxLength(nums []int) int {
	n := len(nums)
	res := 0
	for i := n - 1; i >= 0; i-- {
		p := big.NewInt(1)
		g := 0
		l := 1
		for j := i; j < n; j++ {
			p.Mul(p, big.NewInt(int64(nums[j])))
			if g == 0 {
				g = nums[j]
			} else {
				g = gcd(g, nums[j])
			}
			l = lcm(l, nums[j])
			if p.Cmp(big.NewInt(int64(g*l))) == 0 {
				if j-i+1 > res {
					res = j - i + 1
				}
			}
		}
	}
	return res
}
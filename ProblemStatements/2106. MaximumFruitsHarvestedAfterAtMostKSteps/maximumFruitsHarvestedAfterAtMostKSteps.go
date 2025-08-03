func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + fruits[i][1]
	}
	res, l := 0, 0
	for r := 0; r < n; r++ {
		for l <= r {
			l1 := fruits[l][0]
			r1 := fruits[r][0]
			h := min(abs(startPos-l1)+r1-l1, abs(startPos-r1)+r1-l1)
			if h <= k {
				break
			}
			l++
		}
		if l <= r {
			sum := pre[r+1] - pre[l]
			if sum > res {
				res = sum
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
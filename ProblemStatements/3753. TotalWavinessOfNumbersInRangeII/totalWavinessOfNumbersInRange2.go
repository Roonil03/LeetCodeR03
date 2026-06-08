func totalWaviness(num1 int64, num2 int64) int64 {
	return countWaviness(num2) - countWaviness(num1-1)
}

func countWaviness(n int64) int64 {
	if n < 100 {
		return 0
	}

	var ans int64

	for pattern := 0; pattern < 1000; pattern++ {
		right := pattern % 10
		mid := pattern / 10 % 10
		left := pattern / 100

		if !isWaveTriple(left, mid, right) {
			continue
		}

		minPrefix := int64(0)

		if left == 0 {
			minPrefix = 1
		}

		for mult := int64(1); mult*100 <= n; mult *= 10 {
			pow := mult * 1000

			prefix := n / pow
			current := (n / mult) % 1000
			suffix := n % mult

			p := int64(pattern)

			if current > p {
				if prefix >= minPrefix {
					ans += (prefix - minPrefix + 1) * mult
				}
			} else if current == p {
				if prefix >= minPrefix {
					ans += (prefix - minPrefix) * mult
					ans += suffix + 1
				}
			} else {
				if prefix > minPrefix {
					ans += (prefix - minPrefix) * mult
				}
			}
		}
	}
	return ans
}

func isWaveTriple(left, mid, right int) bool {
	return (mid > left && mid > right) || (mid < left && mid < right)
}
func countAnagrams(s string) int {
	const MOD = 1000000007
	words := strings.Fields(s)
	res := 1
	for _, word := range words {
		freq := make(map[byte]int)
		for i := 0; i < len(word); i++ {
			freq[word[i]]++
		}
		curr := factorial(len(word))
		for _, count := range freq {
			if count > 1 {
				curr = (curr * modInverse(factorial(count), MOD)) % MOD
			}
		}
		res = (res * curr) % MOD
	}
	return res
}

func factorial(n int) int {
	const MOD = 1000000007
	res := 1
	for i := 2; i <= n; i++ {
		res = (res * i) % MOD
	}
	return res
}

func modInverse(a, m int) int {
	return modPow(a, m-2, m)
}

func modPow(x, n, m int) int {
	if n == 0 {
		return 1
	}
	temp := modPow(x, n/2, m)
	temp = (temp * temp) % m
	if n%2 == 1 {
		temp = (temp * x) % m
	}
	return temp
}
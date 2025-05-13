func lengthAfterTransformations(s string, t int) int {
	const mod int = 1e9 + 7
	count := make([]int, 26)
	for _, char := range s {
		count[char-'a']++
	}
	for j := 0; j < t; j++ {
		temp := make([]int, 26)
		for i := 0; i < 26; i++ {
			if i == 25 {
				temp[0] = (temp[0] + count[i]) % mod
				temp[1] = (temp[1] + count[i]) % mod
			} else {
				temp[i+1] = (temp[i+1] + count[i]) % mod
			}
		}
		count = temp
	}
	res := 0
	for _, c := range count {
		res = (res + c) % mod
	}
	return res
}
func kthFactor(n int, k int) int {
	// var fact []int
	// sqrtN := int(math.Sqrt(float64(n)))
	// for i := 1; i <= sqrtN; i++ {
	//     if n%i == 0 {
	//         k--
	//         fact = append(fact, i)
	//         if k == 0 {
	//             return i
	//         }
	//     }
	// }
	// if sqrtN*sqrtN == n {
	//     sqrtN--
	// }
	// for i := len(fact) - 1; i >= 0; i-- {
	//     if n/fact[i] != fact[i] {
	//         k--
	//         if k == 0 {
	//             return n / fact[i]
	//         }
	//     }
	// }
	// return -1
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			k -= 1
			if k == 0 {
				return i
			}
		}
	}
	return -1
}
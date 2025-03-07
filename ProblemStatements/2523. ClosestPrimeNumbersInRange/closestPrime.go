func closestPrimes(left int, right int) []int {
	isPrime := sieveOfEratosthenes(right)
	prevPrime := -1
	closestPair := []int{-1, -1}
	smallestDiff := math.MaxInt32
	for num := left; num <= right; num++ {
		if isPrime[num] {
			if prevPrime != -1 && num-prevPrime < smallestDiff {
				smallestDiff = num - prevPrime
				closestPair[0] = prevPrime
				closestPair[1] = num
			}
			prevPrime = num
		}
	}
	return closestPair
}

func sieveOfEratosthenes(max int) []bool {
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= max; p++ {
		if isPrime[p] {
			for i := p * p; i <= max; i += p {
				isPrime[i] = false
			}
		}
	}
	return isPrime
}

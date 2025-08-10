func reorderedPowerOf2(n int) bool {
	m := make(map[string]bool, 30)
	for i := 1; i < 1000000000; i <<= 1 {
		m[h1(i)] = true
	}
	_, ok := m[h1(n)]
	return ok
}

func h1(a int) string {
	count := [10]byte{}
	for a > 0 {
		count[a%10]++
		a /= 10
	}
	return string(count[:])
}
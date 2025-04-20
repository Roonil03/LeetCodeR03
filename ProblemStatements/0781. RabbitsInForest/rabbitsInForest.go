func numRabbits(answers []int) int {
	m := make(map[int]int)
	res := 0
	for _, i := range answers {
		m[i]++
	}
	for k, v := range m {
		res += int(math.Ceil(float64(v)/float64(k+1)) * float64(k+1))
	}
	return res
}
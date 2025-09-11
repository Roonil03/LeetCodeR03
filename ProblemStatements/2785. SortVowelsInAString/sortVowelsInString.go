func sortVowels(s string) string {
	res := []byte(s)
	vow := make([]byte, 0)
	for i := 0; i < len(res); i++ {
		switch res[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vow = append(vow, res[i])
		}
	}
	sort.Slice(vow, func(i, j int) bool {
		return vow[i] < vow[j]
	})
	y := 0
	for i := 0; i < len(res); i++ {
		switch res[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			res[i] = vow[y]
			y++
		}
	}
	return string(res)
}
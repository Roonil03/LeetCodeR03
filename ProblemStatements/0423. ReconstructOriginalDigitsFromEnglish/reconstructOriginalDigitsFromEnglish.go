func originalDigits(s string) string {
	freq := make([]int, 10)
	for _, i := range s {
		switch i {
		case 'z':
			freq[0]++
		case 'w':
			freq[2]++
		case 'x':
			freq[6]++
		case 'g':
			freq[8]++
		case 'u':
			freq[4]++
		case 's':
			freq[7]++
		case 'f':
			freq[5]++
		case 'h':
			freq[3]++
		case 'i':
			freq[9]++
		case 'o':
			freq[1]++
		}
	}
	freq[7] -= freq[6]
	freq[5] -= freq[4]
	freq[3] -= freq[8]
	freq[9] = freq[9] - freq[8] - freq[5] - freq[6]
	freq[1] = freq[1] - freq[0] - freq[2] - freq[4]
	res := bytes.Buffer{}
	for i := range freq {
		for j := 0; j < freq[i]; j++ {
			_ = res.WriteByte('0' + byte(i))
		}
	}
	return res.String()
}
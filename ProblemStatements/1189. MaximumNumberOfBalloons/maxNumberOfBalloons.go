func maxNumberOfBalloons(text string) int {
	freq := make([]int, 5)
	for i := range text {
		if text[i] == 'b' {
			freq[0]++
		} else if text[i] == 'a' {
			freq[1]++
		} else if text[i] == 'l' {
			freq[2]++
		} else if text[i] == 'o' {
			freq[3]++
		} else if text[i] == 'n' {
			freq[4]++
		}
	}
	return min(freq[0], min(freq[1], min(freq[2]/2, min(freq[3]/2, freq[4]))))
}
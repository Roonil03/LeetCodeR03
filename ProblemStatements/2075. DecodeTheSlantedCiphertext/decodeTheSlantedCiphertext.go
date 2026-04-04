func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	numCols := n / rows
	if n == 0 {
		return ""
	}
	decodedText := make([]rune, n)
	var i, j, k, lastIdx int
	for {
		idx := i*numCols + j
		if idx >= n {
			break
		}
		decodedText[k] = rune(encodedText[i*numCols+j])
		i++
		j++
		k++
		if i >= rows {
			i -= rows
			j -= rows
			j++
		}
		if string(encodedText[idx]) != " " {
			lastIdx = k
		}
	}
	return string(decodedText[:lastIdx])
}
func minimumDeletions(word string, k int) int {
	freq := make(map[rune]int)
	for _, c := range word {
		freq[c]++
	}
	counts := []int{}
	for _, v := range freq {
		counts = append(counts, v)
	}
	res := math.MaxInt32
	n := len(counts)
	if n == 1 {
		return 0
	}
	for minFreq := 1; minFreq <= len(word); minFreq++ {
		maxFreq := minFreq + k
		curDel := 0
		for _, f := range counts {
			if f < minFreq {
				curDel += f
			} else if f > maxFreq {
				curDel += f - maxFreq
			}
		}
		if curDel < res {
			res = curDel
		}
	}
	return res
}
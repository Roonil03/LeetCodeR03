// /**
//  * // This is the Master's API interface.
//  * // You should not implement it, or speculate about its implementation
//  * type Master struct {
//  * }
//  *
//  * func (this *Master) Guess(word string) int {}
//  */
// func findSecretWord(words []string, master *Master) {
//     for i := 0; i < 10; i++{
//         guess := h0(words)
//         matches := master.Guess(guess)
//         if matches == 6{
//             return
//         }
//         filt := []string{}
//         for _, w := range words{
//             if h1(guess, w) == matches{
//                 filt = append(filt, w)
//             }
//         }
//         words = filt
//     }
// }

// func h0(words []string) string{
//     bestGuess := words[0]
//     minMaxRemaining := math.MaxInt32
//     for _, candidate := range words {
//         matchGroups := make([]int, 7)
//         for _, word := range words {
//             if candidate == word {
//                 continue
//             }
//             matches := h1(candidate, word)
//             matchGroups[matches]++
//         }
//         maxGroupSize := 0
//         for _, groupSize := range matchGroups {
//             if groupSize > maxGroupSize {
//                 maxGroupSize = groupSize
//             }
//         }
//         if maxGroupSize < minMaxRemaining {
//             minMaxRemaining = maxGroupSize
//             bestGuess = candidate
//         }
//     }
//     return bestGuess
// }

// func h1(w1, w2 string) int{
//     res := 0
//     for i := 0; i < len(w1); i++{
//         if w1[i] == w2[i]{
//             res++
//         }
//     }
//     return res
// }

func findSecretWord(words []string, master *Master) {
	index := findUnique(words)
	res := master.Guess(words[index])
	if res == 6 {
		return
	}
	findSimilarWords(words, words[index], res)
	words[index] = ""
	for _, word := range words {
		if len(word) != 6 {
			continue
		}
		res := master.Guess(word)
		if res == 6 {
			return
		}
		findSimilarWords(words, word, res)
	}
}

func findSimilarWords(words []string, word string, matches int) {
	for k, w := range words {
		if len(w) != 6 {
			continue
		}
		count := 0
		for i := 0; i < 6; i++ {
			if w[i] == word[i] {
				count++
			}
		}
		if count != matches {
			words[k] = ""
		}
	}
}

func findUnique(words []string) int {
	index := 0
	mincount := math.MaxInt32
	for i := 0; i < len(words); i++ {
		wordcount := 0
		for j := i + 1; j < len(words); j++ {
			for k := 0; k < 6; k++ {
				if words[i][k] == words[j][k] {
					wordcount++
				}
			}
		}
		if wordcount < mincount {
			mincount = wordcount
			index = i
		}
	}
	return index
}
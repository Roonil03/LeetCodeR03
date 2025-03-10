// func countOfSubstrings(word string, k int) int64 {
//     arr := make([]int, 6)
//     n := len(word)
//     var res int64
//     res = 0
//     for i:= 0; i < (k+5); i++{
//         switch word[i]{
//             case 'a':
//             arr[0]++
//             case 'e':
//             arr[1]++
//             case 'i':
//             arr[2]++
//             case 'o':
//             arr[3]++
//             case 'u':
//             arr[4]++
//             default:
//             arr[5]++
//         }
//     }
//     if isValid(arr, k){
//         res = res + 1
//     }
//     for i:= k+5; i < n; i++{
//         switch word[i-(k+5)]{
//             case 'a':
//             arr[0]--
//             case 'e':
//             arr[1]--
//             case 'i':
//             arr[2]--
//             case 'o':
//             arr[3]--
//             case 'u':
//             arr[4]--
//             default:
//             arr[5]--
//         }
//         switch word[i]{
//             case 'a':
//             arr[0]++
//             case 'e':
//             arr[1]++
//             case 'i':
//             arr[2]++
//             case 'o':
//             arr[3]++
//             case 'u':
//             arr[4]++
//             default:
//             arr[5]++
//         }
//         if isValid(arr, k){
//             res = res + 1
//         }
//     }
//     return res
// }

// func isValid(arr []int, k int) bool{
//     for i := 0; i < len(arr)-1; i++{
//         if arr[i] < 1{
//             return false
//         }
//     }
//     return arr[len(arr)-1] == k
// }

// func isVowel(ch byte) bool {
// 	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
// }

// func isValidWindow(vowelCount map[byte]int, consonantCount int, k int) bool {
// 	if vowelCount['a'] == 0 || vowelCount['e'] == 0 || vowelCount['i'] == 0 ||
// 	   vowelCount['o'] == 0 || vowelCount['u'] == 0 {
// 		return false
// 	}
// 	return consonantCount == k
// }

// func countOfSubstrings(word string, k int) int64 {
// 	n := len(word)
// 	if n < k+5 {
// 		return 0
// 	}
// 	var res int64 = 0
// 	left, right := 0, 0
// 	vowelCount := make(map[byte]int)
// 	consonantCount := 0
// 	minSize := k + 5
// 	maxSize := n
// 	for size := minSize; size <= maxSize; size++ {
// 		vowelCount = make(map[byte]int)
// 		consonantCount = 0
// 		left, right = 0, 0
// 		for right < size {
// 			if right >= n {
// 				break
// 			}
// 			if isVowel(word[right]) {
// 				vowelCount[word[right]]++
// 			} else {
// 				consonantCount++
// 			}
// 			right++
// 		}
// 		if isValidWindow(vowelCount, consonantCount, k) {
// 			res++
// 		}
// 		for right < n {
// 			if isVowel(word[left]) {
// 				vowelCount[word[left]]--
// 			} else {
// 				consonantCount--
// 			}
// 			left++
// 			if isVowel(word[right]) {
// 				vowelCount[word[right]]++
// 			} else {
// 				consonantCount++
// 			}
// 			right++
// 			if isValidWindow(vowelCount, consonantCount, k) {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

// func isVowel(ch byte) bool {
// 	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
// }

// func atLeastK(word string, k int) int64 {
// 	vowelCount := make(map[byte]int)
// 	var consonantCount, start, end int
// 	var numValidSubStrings int64
// 	for end < len(word) {
// 		newLetter := word[end]
// 		if isVowel(newLetter) {
// 			vowelCount[newLetter]++
// 		}else {
// 			consonantCount++
// 		}
// 		for len(vowelCount) == 5 && consonantCount >= k {
// 			numValidSubStrings += int64(len(word) - end)
// 			startLetter := word[start]
// 			if isVowel(startLetter) {
// 				vowelCount[startLetter]--
// 				if vowelCount[startLetter] == 0 {
// 					delete(vowelCount, startLetter)
// 				}
// 			} else {
// 				consonantCount--
// 			}
// 			start++
// 		}
// 		end++
// 	}
// 	return numValidSubStrings
// }

// func countOfSubstrings(word string, k int) int64 {
// 	return atLeastK(word, k) - atLeastK(word, k + 1)
// }

func countOfSubstrings(word string, k int) int64 {
	a := make([]int, 5)
	a[0], a[1], a[2], a[3], a[4] = -1, -1, -1, -1, -1
	n := 0
	n1, n2 := 0, 0
	s1, s2 := -1, -1

	var ret int64 = 0
	for i := 0; i < len(word); i++ {
		idx := vowel(word[i])
		if idx != -1 {
			a[idx] = i
		} else {
			n++
		}

		if n < k {
			continue
		}

		right := min(a[0], min(a[1], min(a[2], min(a[3], a[4]))))
		if right == -1 {
			continue
		}
		for s1 < i && n1 < n-k {
			s1++
			if vowel(word[s1]) == -1 {
				n1++
			}
		}

		for s2 < i && n2 < n-k+1 {
			s2++
			if vowel(word[s2]) == -1 {
				n2++
			}
		}

		tmp := max(min(right, s2)-s1, 0)
		ret += int64(tmp)
	}

	return ret
}

func vowel(a byte) int {
	if a == 'a' {
		return 0
	} else if a == 'e' {
		return 1
	} else if a == 'i' {
		return 2
	} else if a == 'o' {
		return 3
	} else if a == 'u' {
		return 4
	} else {
		return -1
	}
}
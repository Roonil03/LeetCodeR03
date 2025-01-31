// func reverseWords(s string) string {
//     w := strings.Split(s, " ")
//     for i, j:= 0, 0; i < (len(w)/2); i++{
//         if w[i] == " "{
//             j++
//         }
//         temp := w[i+j]
//         w[i] = w[len(w)-i-j-1]
//         w[len(w) - i - j- 1] = temp
//     }
//     return strings.Join(w, " ")
// }

func reverseWords(s string) string {
	w := []byte(s)
	length := removeExtraSpaces(w)
	w = w[:length]
	reverse(w, 0, len(w)-1)
	t1 := 0
	for i := 0; i < len(w); i++ {
		if i == len(w) || w[i] == ' ' {
			reverse(w, t1, i-1)
			t1 = i + 1
		}
	}
	reverse(w, t1, len(w)-1)
	return string(w)
}

func removeExtraSpaces(w []byte) int {
	i := 0
	j := 0
	for i < len(w) && w[i] == ' ' {
		i++
	}
	for i < len(w) {
		w[j] = w[i]
		j++
		i++
		for i < len(w) && w[i] == ' ' {
			if w[i-1] != ' ' {
				w[j] = ' '
				j++
			}
			i++
		}
	}
	if j > 0 && w[j-1] == ' ' {
		j--
	}
	return j
}

func reverse(w []byte, t1, t2 int) {
	for t1 < t2 {
		w[t1], w[t2] = w[t2], w[t1]
		t1++
		t2--
	}
}
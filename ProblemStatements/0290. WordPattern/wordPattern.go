// func wordPattern(pattern string, s string) bool {
//     w := strings.Split(s, " ")
//     if len(w) != len(pattern) {
//         return false
//     }
//     idx := make(map[interface{}]int)
//     for i := 0; i < len(w); i++ {
//         pat := string(pattern[i])
//         key := w[i]
//         pv1, ex1 := idx[pat]
//         pv2, ex2 := idx[key]
//         if (ex1 != ex2) || (ex1 && ex2 && pv1 != pv2) {
//             return false
//         }
//         idx[pat] = i
//         idx[key] = i
//     }
//     return true
// }

func wordPattern(pattern string, str string) bool {
	codes := make(map[string]byte)
	used := make([]bool, 26)
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		code, ok := codes[word]
		if !ok {
			if used[pattern[i]-'a'] {
				return false
			}
			used[pattern[i]-'a'] = true
			codes[word] = pattern[i]
			continue
		}
		if code != pattern[i] {
			return false
		}
	}
	return true
}
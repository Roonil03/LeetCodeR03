func possibleStringCount(word string) int {
	// n := len(word)
	// if n == 0 {
	//     return 0
	// }
	// res := 0
	// for i:= 0; i < n;{
	//     j := 1
	//     for j + 1 < n && word[j + 1] == word[i]{
	//         j++
	//     }
	//     if j > i {
	//         // p := word[:i + 1] + word[j + 1:]
	//         res++
	//     }
	//     i = j + 1
	// }
	// res++
	// return res
	res := 1
	ct := 1
	ch := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == ch {
			ct++
		} else {
			res += (ct - 1)
			ch = word[i]
			ct = 1
		}
	}
	res += (ct - 1)
	return res
}
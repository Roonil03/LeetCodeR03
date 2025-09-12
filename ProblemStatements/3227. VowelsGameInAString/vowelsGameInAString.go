func doesAliceWin(s string) bool {
	// vow := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
	// for i := 0; i < len(s); i++{
	//     if vow[s[i]]{
	//         return true
	//     }
	// }
	// return false
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			return true
		}
	}
	return false
}
func isSubsequence(s string, t string) bool {
	// i :=  0
	// m, n := len(s), len(t)
	// if m > n{
	//     return false
	// }
	// if m == 0{
	//     return true
	// }
	// for a := 0 ; a < n; a++{
	//     if i == m-1{
	//         return true
	//     }
	//     if s[i] == t[a]{
	//         i++
	//     }
	// }
	// return false
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

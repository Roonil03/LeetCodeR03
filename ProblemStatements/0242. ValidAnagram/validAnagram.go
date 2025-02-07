// func isAnagram(s string, t string) bool {
//     if len(s) != len(t){
//         return false
//     }
//     a := make([]int, 26)
//     n := len(s)
//     for i:=0;i<n;i++{
//         a[s[i]-'a']++
//     }
//     for i :=0;i < n; i++{
//         a[t[i]-'a']--
//         if a[t[i]-'a'] == -1{
//             return false
//         }
//     }
//     return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	let := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		let[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if let[ransomNote[i]-'a'] == 0 {
			return false
		}
		let[ransomNote[i]-'a']--
	}
	return true
}
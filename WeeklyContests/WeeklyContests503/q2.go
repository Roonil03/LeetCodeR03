func passwordStrength(password string) int {
	var vis [256]bool
	res := 0
	for i := range password {
		b := password[i]
		if vis[b] {
			continue
		}
		vis[b] = true
		switch {
		case b >= 'a' && b <= 'z':
			res += 1
		case b >= 'A' && b <= 'Z':
			res += 2
		case b >= '0' && b <= '9':
			res += 3
		case b == '!' || b == '@' || b == '#' || b == '$':
			res += 5
		}
	}
	return res
}
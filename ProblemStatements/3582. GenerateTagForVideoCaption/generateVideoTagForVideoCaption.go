func generateTag(caption string) string {
	a := strings.Fields(caption)
	if len(a) == 0 {
		return "#"
	}
	var res strings.Builder
	res.WriteByte('#')
	for i, w := range a {
		for j, r := range w {
			if unicode.IsLetter(r) && r <= 127 {
				if i == 0 {
					res.WriteRune(unicode.ToLower(r))
				} else {
					if j == 0 {
						res.WriteRune(unicode.ToUpper(r))
					} else {
						res.WriteRune(unicode.ToLower(r))
					}
				}
			}
			if res.Len() >= 100 {
				return res.String()
			}
		}
	}
	return res.String()
}
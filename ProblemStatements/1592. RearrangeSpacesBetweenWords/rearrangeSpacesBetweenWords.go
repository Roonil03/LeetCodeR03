// import "strings"

func reorderSpaces(text string) string {
	sc := 0
	for _, c := range text {
		if c == ' ' {
			sc++
		}
	}
	words := strings.Fields(text)
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", sc)
	}
	gaps := len(words) - 1
	sb := sc / gaps
	e := sc % gaps
	res := strings.Join(words, strings.Repeat(" ", sb))
	res += strings.Repeat(" ", e)
	return res
}
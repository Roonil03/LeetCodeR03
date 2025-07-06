func concatHex36(n int) string {
	sq := int64(n) * int64(n)
	c := int64(n) * sq
	h1 := strings.ToUpper(strconv.FormatInt(sq, 16))
	h2 := strings.ToUpper(strconv.FormatInt(c, 36))
	return h1 + h2
}
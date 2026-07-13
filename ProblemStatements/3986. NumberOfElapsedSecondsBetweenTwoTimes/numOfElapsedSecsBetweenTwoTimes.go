func secondsBetweenTimes(st string, et string) int {
	t1 := int(st[0])*36000 + int(st[1])*3600 + int(st[3])*600 + int(st[4])*60 + int(st[6])*10 + int(st[7])
	t2 := int(et[0])*36000 + int(et[1])*3600 + int(et[3])*600 + int(et[4])*60 + int(et[6])*10 + int(et[7])
	res := (t2 - t1 + 86400) % 86400
	return res
}
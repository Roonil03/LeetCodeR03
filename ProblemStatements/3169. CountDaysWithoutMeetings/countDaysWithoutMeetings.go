func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	a := 0
	for _, m := range meetings {
		s := max(m[0], a+1)
		l := m[1] - s + 1
		days -= max(l, 0)
		a = max(a, m[1])
	}
	return days
}
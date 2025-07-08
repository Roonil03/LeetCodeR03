func maxValue(e [][]int, k int) int {
	type Event struct {
		start int
		end   int
		value int
	}
	n := len(e)
	events := make([]Event, n)
	for i, e1 := range e {
		events[i] = Event{
			start: e1[0],
			end:   e1[1],
			value: e1[2],
		}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].start < events[j].start
	})
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(int, int) int
	dp = func(id int, rem int) int {
		if id == n || rem == 0 {
			return 0
		}
		if memo[id][rem] != -1 {
			return memo[id][rem]
		}
		skip := dp(id+1, rem)
		var h1 func([]Event, int) int
		h1 = func(h []Event, endDay int) int {
			l, r := 0, len(h)
			for l < r {
				m := l + (r-l)/2
				if h[m].start > endDay {
					r = m
				} else {
					l = m + 1
				}
			}
			return l
		}
		ni := h1(events, events[id].end)
		take := events[id].value + dp(ni, rem-1)
		res := max(skip, take)
		memo[id][rem] = res
		return res
	}
	return dp(0, k)
}
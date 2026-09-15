package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	r1 := make([]int, len(s2)+1)
	r2 := make([]int, len(s2)+1)
	c, id := 0, 0
	for i := 1; i <= n1; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[id] {
				id++
				if id == len(s2) {
					id = 0
					c++
				}
			}
		}
		if r1[id] > 0 {
			rem := (n1 - i) / (i - r1[id])
			c += rem * (c - r2[id])
			i += rem * (i - r1[id])
			r1 = make([]int, len(s2)+1)
		} else {
			r1[id] = i
			r2[id] = c
		}
	}
	return c / n2
}

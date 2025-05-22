func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	res := make([]bool, n)
	a := candies[0]
	for _, c := range candies {
		if c > a {
			a = c
		}
	}
	for i, c := range candies {
		if c+extraCandies >= a {
			res[i] = true
		}
	}
	return res
}
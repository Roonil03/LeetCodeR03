func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	res := 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			res++
			j++
			i++
		} else {
			j++
		}
	}
	return res
}
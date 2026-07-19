func canReach(start []int, target []int) bool {
	return (start[0]+start[1]+target[0]+target[1])%2 == 0
}
func winningPlayer(x int, y int) string {
	if x < int(y/4) {
		if x%2 == 0 {
			return "Bob"
		} else {
			return "Alice"
		}
	} else {
		if int(y/4)%2 == 0 {
			return "Bob"
		} else {
			return "Alice"
		}
	}

}
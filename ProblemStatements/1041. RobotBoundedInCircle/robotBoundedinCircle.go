func isRobotBounded(instructions string) bool {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y := 0, 0
	dir := 0
	for _, i := range instructions {
		switch i {
		case 'G':
			x += dirs[dir][0]
			y += dirs[dir][1]
		case 'L':
			dir = (dir + 3) % 4
		case 'R':
			dir = (dir + 1) % 4
		}
	}
	return (x == 0 && y == 0) || (dir != 0)
}
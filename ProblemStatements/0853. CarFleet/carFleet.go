func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	car := make([]Car, n)
	for i := 0; i < n; i++ {
		car[i] = Car{
			pos:   position[i],
			speed: speed[i],
			time:  float64(target-position[i]) / float64(speed[i]),
		}
	}
	sort.Slice(car, func(i, j int) bool {
		return car[i].pos > car[j].pos
	})
	res := 0
	temp := 0.0
	for _, c := range car {
		if c.time > temp {
			res++
			temp = c.time
		}
	}
	return res
}

type Car struct {
	pos   int
	speed int
	time  float64
}
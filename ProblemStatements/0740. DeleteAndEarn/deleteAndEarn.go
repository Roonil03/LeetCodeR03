func deleteAndEarn(nums []int) int {
	points := make(map[int]int)
	maxNum := 0
	for _, num := range nums {
		points[num] += num
		if num > maxNum {
			maxNum = num
		}
	}
	unique := make([]int, 0, len(points))
	for num := range points {
		unique = append(unique, num)
	}
	sort.Ints(unique)
	prev := -1
	take, skip := 0, 0
	for _, num := range unique {
		currentPoints := points[num]
		if num == prev+1 {
			take, skip = skip+currentPoints, max(take, skip)
		} else {
			take, skip = max(take, skip)+currentPoints, max(take, skip)
		}
		prev = num
	}
	return max(take, skip)
}
func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// a,b := slices.Min(landStartTime), slices.Min(waterStartTime)
	// res := math.MaxInt
	// for i := range landStartTime{
	//     x := max(landStartTime[i] + landDuration[i], b) + waterDuration[slices.Index(waterStartTime, b)]
	//     res = min(res, x)
	// }
	// for i := range waterStartTime{
	//     x := max(waterStartTime[i] + waterDuration[i], a) + landDuration[slices.Index(landStartTime, a)]
	//     res = min(res, x)
	// }
	// return res

	res := h1(landStartTime, landDuration, waterStartTime, waterDuration)
	res = min(res, h1(waterStartTime, waterDuration, landStartTime, landDuration))
	return res
}

func h1(l1StartTime, l1Duration, l2StartTime, l2Duration []int) int {
	l1MinEndTime := math.MaxInt
	for i := range l1StartTime {
		l1MinEndTime = min(l1MinEndTime, l1StartTime[i]+l1Duration[i])
	}
	l2MinEndTime := math.MaxInt
	for i := range l2StartTime {
		if l1MinEndTime < l2StartTime[i] {
			l2MinEndTime = min(l2MinEndTime, l2StartTime[i]+l2Duration[i])
		} else {
			l2MinEndTime = min(l2MinEndTime, l1MinEndTime+l2Duration[i])
		}
	}
	return l2MinEndTime
}
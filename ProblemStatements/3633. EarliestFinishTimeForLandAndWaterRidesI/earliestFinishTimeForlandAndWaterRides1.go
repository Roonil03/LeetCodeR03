func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	minFL := int(^uint(0) >> 1)
	for i := 0; i < len(landStartTime); i++ {
		fl := landStartTime[i] + landDuration[i]
		if fl < minFL {
			minFL = fl
		}
	}
	minFW := int(^uint(0) >> 1)
	for j := 0; j < len(waterStartTime); j++ {
		fw := waterStartTime[j] + waterDuration[j]
		if fw < minFW {
			minFW = fw
		}
	}
	res := int(^uint(0) >> 1)
	for j := 0; j < len(waterStartTime); j++ {
		start := minFL
		if waterStartTime[j] > start {
			start = waterStartTime[j]
		}
		fin := start + waterDuration[j]
		if fin < res {
			res = fin
		}
	}
	for i := 0; i < len(landStartTime); i++ {
		start := minFW
		if landStartTime[i] > start {
			start = landStartTime[i]
		}
		fin := start + landDuration[i]
		if fin < res {
			res = fin
		}
	}
	return res
}
func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	pairs := make([]Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = Pair{nums[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].element < pairs[j].element
	})
	result := [][]Pair{}
	currentGroup := []Pair{pairs[0]}
	for i := 1; i < n; i++ {
		if pairs[i].element-pairs[i-1].element <= limit {
			currentGroup = append(currentGroup, pairs[i])
		} else {
			result = append(result, currentGroup)
			currentGroup = []Pair{pairs[i]}
		}
	}
	result = append(result, currentGroup)
	for _, group := range result {
		sortedGroup := append([]Pair(nil), group...)
		sort.Slice(sortedGroup, func(i, j int) bool {
			return sortedGroup[i].index < sortedGroup[j].index
		})
		for i, pair := range sortedGroup {
			nums[pair.index] = group[i].element
		}
	}
	return nums
}

type Pair struct {
	element int
	index   int
}
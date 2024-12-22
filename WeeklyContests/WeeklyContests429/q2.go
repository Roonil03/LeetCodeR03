/* Did not submit during the contest to make sure  I don't get much penalty,
also don't understand why this works, but my version doesn't
*/

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	tmp := []int{}
	for _, v := range nums {
		if len(tmp) > 0 {
			if tmp[len(tmp)-1] >= v+k {
				continue
			} else {
				tmp = append(tmp, max(tmp[len(tmp)-1]+1, v-k))
			}
		} else {
			tmp = append(tmp, v-k)
		}
	}
	return len(tmp)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


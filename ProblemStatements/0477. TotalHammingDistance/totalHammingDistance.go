func totalHammingDistance(nums []int) int {
	n, res := len(nums), 0
	for i := 0; i < 32; i++ {
		o := 0
		for _, a := range nums {
			if (a>>i)&1 == 1 {
				o++
			}
		}
		z := n - o
		res += z * o
	}
	return res
}
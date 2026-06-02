func minElement(nums []int) int {
	res := math.MaxInt32
	for i := range nums {
		res = min(h1(nums[i]), res)
	}
	return res
}

func h1(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
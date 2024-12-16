package main

func getSumAbsoluteDifferences(nums []int) []int {
	len := len(nums)
	res := make([]int, len)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	prefix := 0
	for i := 0; i < len; i++ {
		left := nums[i]*i - prefix
		right := (sum - prefix - nums[i]) - nums[i]*(len-i-1)
		res[i] = left + right
		prefix += nums[i]
	}
	return res
}

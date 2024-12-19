package main

func canJump(nums []int) bool {
	a := 0
	for i := 0; i < len(nums); i++ {
		if i > a {
			return false
		}
		a = max(a, i+nums[i])
		if a >= len(nums)-1 {
			return true
		}
	}
	return false
}

func triangleType(nums []int) string {
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]
	if a+b <= c {
		return "none"
	}
	if a == c {
		return "equilateral"
	}
	if a == b || b == c {
		return "isosceles"
	}
	return "scalene"
}
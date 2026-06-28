func maxSubarraySum(nums []int, k int) int64 {
	var dp0m, dp1m, dp2m int64 = 0, -1e16, -1e16
	var dp0d, dp1d, dp2d int64 = 0, -1e16, -1e16
	var res int64 = -1e16
	for _, v := range nums {
		val := int64(v)
		mv := val * int64(k)
		dv := val / int64(k)
		nd2m := max(dp1m+val, dp2m+val)
		nd1m := max(mv, dp0m+mv, dp1m+mv)
		dp0m = max(val, dp0m+val)
		dp1m, dp2m = nd1m, nd2m
		nd2d := max(dp1d+val, dp2d+val)
		nd1d := max(dv, dp0d+dv, dp1d+dv)
		dp0d = max(val, dp0d+val)
		dp1d, dp2d = nd1d, nd2d
		res = max(res, dp1m, dp2m, dp1d, dp2d)
	}
	return res
}
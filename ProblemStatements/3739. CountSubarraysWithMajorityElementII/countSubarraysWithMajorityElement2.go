func countMajoritySubarrays(nums []int, target int) int64 {
    n := len(nums)
    o := n
    counts := make([]int, 2 * n + 1)
    counts[0+o] = 1
    var res int64 = 0
    pre := 0
    lcur := 0
    for _, v := range nums{
        if v == target{
            lcur += counts[pre + o]
            pre++
        } else{
            pre--
            lcur -= counts[pre + o]
        }
        res += int64(lcur)
        counts[pre + o]++
    }
    return res
}
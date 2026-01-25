func minimumDifference(nums []int, k int) int {
    if k == 1{
        return 0
    }
    sort.Ints(nums)
    res := math.MaxInt
    for i := 0; i + k - 1 < len(nums); i++{
        temp := nums[i + k - 1] - nums[i]
        if temp < res{
            res = temp
        }
    }
    return res
}
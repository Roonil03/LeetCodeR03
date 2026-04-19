func maxDistance(nums1 []int, nums2 []int) int {
	var i, j int
	for i, j = 0, 0; i < len(nums1) && j < len(nums2); j++ {
		if nums1[i] > nums2[j] {
			i += 1
		}
	}
	return max(0, j-i-1)
}
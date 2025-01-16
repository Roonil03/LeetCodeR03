class Solution {
    public int xorAllNums(int[] nums1, int[] nums2) {
        int x = 0, y = 0, res = 0;
        for(int n: nums1)
        {
            x ^= n;
        }
        for(int n : nums2)
        {
            y ^= n;
        }
        if(nums2.length %2 != 0)
        {
            res ^= x;
        }
        if(nums1.length % 2 != 0)
        {
            res ^= y;
        }
        return res;
    }
}
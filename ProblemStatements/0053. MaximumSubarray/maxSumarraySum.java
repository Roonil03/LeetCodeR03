class Solution {
    public int maxSubArray(int[] nums) {
        return maxSum(nums, 0, nums.length - 1);
    }

    int maxSum(int[] nums, int l, int r) {
        if (l == r) {
            return nums[l];
        }
        int mid = l + (r - l) / 2;
        int left = maxSum(nums, l, mid);
        int right = maxSum(nums, mid + 1, r);
        int cross = maxCross(nums, l, mid, r);
        return Math.max(Math.max(left, right), cross);
    }

    int maxCross(int[] nums, int l, int m, int r) {
        int lsum = Integer.MIN_VALUE;
        int sum = 0;
        for (int i = m; i >= l; i--) {
            sum += nums[i];
            lsum = Math.max(lsum, sum);
        }
        int rsum = Integer.MIN_VALUE;
        sum = 0;
        for (int i = m + 1; i <= r; i++) {
            sum += nums[i];
            rsum = Math.max(rsum, sum);
        }
        return lsum + rsum;
    }
}

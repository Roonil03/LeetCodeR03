class Solution {
    public int countSubarrays(int[] nums) {
        int count = 0;
        for(int i = 0; i< nums.length - 2; i++)
            {
                if((nums[i] + nums[i+2]) == (nums[i+1]/(float)2))
                    count++;
            }
        return count;
    }
}
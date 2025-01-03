public class Solution {
    public int WaysToSplitArray(int[] nums) {
        long r = 0L;
        foreach(int n in nums){
            r += n;
        }
        long l = 0L;
        int ans = 0;
        for(int i = 0; i< nums.Length - 1; i++){
            l += nums[i];
            r -= nums[i];
            if(l >= r){
                ans++;
            }
        }
        return ans;
    }
}
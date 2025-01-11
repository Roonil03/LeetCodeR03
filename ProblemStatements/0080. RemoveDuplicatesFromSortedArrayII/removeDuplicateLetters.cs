public class Solution {
    public int RemoveDuplicates(int[] nums) {
        int i = 0;
        foreach (int m in nums) {            
            if (i < 2 || m != nums[i - 2]) {
                nums[i] = m;
                i++;
            }
        }
        return i;
    }
}
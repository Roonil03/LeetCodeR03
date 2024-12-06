import java.util.*;

class Solution {
    public int firstMissingPositive(int[] nums) {
        int i = 1;
        Arrays.sort(nums);
        for(int n : nums)
        {
            if(n > 0 && n == i)
            {
                i++;
            }
            else if(n>i)
            {
                return i;
            }
        }
        return i;
    }
} 
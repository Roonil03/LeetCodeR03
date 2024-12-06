import java.util.*;

class Solution {
    public int firstMissingPositive(int[] nums) {
        Hashtable<Integer, Integer> h = new Hashtable<>();
        for(int i = 0; i<nums.length; i++)
        {
            if(nums[i]>0)
            {
                h.put(nums[i],nums[i]);
            }
        }
        int i = 1;
        while(true)
        {
            if(!(h.containsKey(i)))
            {
                return i;
            }
            i++;
        }
        //return -1;
    }
}
public class Solution {
    public void SortColors(int[] nums) {
        int n1 = 0, n2 = nums.Length-1;
        int i = 0;
        while(i <= n2)
        {
            if(nums[i] == 0){
                int temp = nums[i];
                nums[i] = nums[n1];
                nums[n1++] = temp;
                i++;
            }
            else if(nums[i] == 2){
                int temp = nums[i];
                nums[i] = nums[n2];
                nums[n2--] = temp;
            }
            else{
                i++;
            }
        }
    }
}
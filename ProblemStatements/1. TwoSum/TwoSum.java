import java.util.*;
public class TwoSum {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter the number of elements in the array");
        int n = sc.nextInt();
        int nums[] = new int[n];
        System.out.println("Enter the elements in the array: ");
        for(int i = 0; i<n;i++)
        {
            nums[i] = sc.nextInt();
        }
        System.out.println("Enter the target: ");
        int target = sc.nextInt();
        //Arrays.sort(nums);
        try{
            int[] ans = Solution.twoSum(nums, target);
            for(int i = 0; i<ans.length;i++){
                System.out.print(ans[i] + " ");
            }
        }
        catch(RuntimeException e)
        {
            System.out.println("Error in argument present");
            System.exit(0);
        }
        finally{
            sc.close();
        }
    }
}
class Solution {
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> hmap = new HashMap<>();
        for(int i = 0; i<nums.length;i++)
        {
            if(hmap.containsKey(target - nums[i])){
                return new int[] {hmap.get(target-nums[i]),i};
            }
            hmap.put(nums[i],i);
        }
        throw new RuntimeException(); //Throws an error in case there has been nothing that returned
    }
}
/* Time  Complexity using the Hashmaps: O(n)
 * Time  Complexity of the first method: O(n)
 * Total Time Complexity of the entire code: O(n)
 * Space Complexity of the entire code: O(n)
*/
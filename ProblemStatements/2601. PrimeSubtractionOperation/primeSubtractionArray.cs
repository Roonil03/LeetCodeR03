public class Solution {
    public bool PrimeSubOperation(int[] nums) {
        if (check(nums)) {
            return true;
        }
        for (int i = 0; i < nums.Length; i++) {
            int bound = (i == 0) ? nums[0] : nums[i] - nums[i - 1];
            if (bound <= 0) {
                return false;
            }
            int largestPrime = 0;
            for (int j = bound - 1; j >= 2; j--) {
                if (isPrime(j)) {
                    largestPrime = j;
                    break;
                }
            }
            if (largestPrime > 0) {
                nums[i] -= largestPrime;
            }
        }
        return check(nums);
    }
    public bool check(int[] nums)
    {
        int count = 0;
        for (int i = 0; i < nums.Length - 1; i++) {
            if (nums[i] >= nums[i + 1]) {
                count++;
            }
        }
        return count == 0;
    }
    private bool isPrime(int num) {
        if (num < 2) {
            return false;
        }
        int val = (int) Math.Sqrt(num);
        for (int i = 2; i <= val; i++) {
            if (num % i == 0) {
                return false;
            }
        }
        return true;
    }
}
/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

public class Solution {
    public int MinLength(string s, int k) {
        List<int> nums = new List<int>();
        foreach (char c in s) {
            nums.Add(c - '0');
        }        
        int n = nums.Count;
        int f1 = 0, f2 = 0;
        for (int i = 0; i < n; i++) {
            if ((nums[i] ^ (i & 1)) != 0)
            {
                f1++;  
            }
            else{
                f2++;
            }
        }

        if (f1 <= k || f2 <= k) {
            return 1;
        }

        List<int> tmp = new List<int>();
        int cur = nums[0], cnt = 0;
        foreach (int v in nums) {
            if (v == cur) {
                cnt++;
            } else {
                tmp.Add(cnt);
                cur = v;
                cnt = 1;
            }
        }
        tmp.Add(cnt);        
        int l = 2, r = n;
        while (l <= r) {
            int m = (l + r) / 2;
            int c = 0;
            foreach (int x in tmp) {
                c += x / (m + 1);
            }
            if (c > k) l = m + 1;
            else r = m - 1;
        }
        return Math.Max(l, 2);
    }
}


// List<int>:   A list is used to store integers (after converting the characters from the input string), 
//              which allows efficient insertion and traversal. It's used for nums to store the converted digits and tmp 
//              to store the frequency of contiguous segments of the same number.

// XOR (^) and Bitwise Operations:  The program uses the XOR operation (^) to determine if the bit at the current index in 
//                                  the string differs from the expected alternating pattern (even/odd indices).

// Frequency Counting:  The program computes contiguous segments of the same value in the string and stores the length of 
//                      each segment in tmp. This helps in determining how many "blocks" of equal numbers there are.

// Binary Search:   The algorithm employs binary search over the length of segments (from 2 to n) to find the minimal segment 
//                  length that satisfies the condition. The binary search minimizes the segment length such that the total number of  
//                  segments larger than that length is less than or equal to k.

// Greedy Count:    During each iteration of the binary search, the program counts how many segments are large enough (i.e., their length 
//                  exceeds m + 1) and adjusts the binary search bounds accordingly.
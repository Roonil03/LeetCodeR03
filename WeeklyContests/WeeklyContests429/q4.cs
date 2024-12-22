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
            if ((nums[i] ^ (i & 1)) != 0) f1++;
            else f2++;
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

// This algorithm solves the problem by first converting the string s into a list of integers and counting how many 
// positions in the string have alternating values based on even and odd indices. It then checks if either of the 
// counts (f1 or f2) is less than or equal to k and returns 1 if true. If not, the algorithm calculates the frequency 
// of contiguous segments of the same number, then uses binary search to minimize the length of segments such that 
// the total number of segments of that length is less than or equal to k. The time complexity is O(n log n) due 
// to the binary search.


// Difference in Algorithms:
    // The first algorithm (from the previous conversation) is primarily focused on segmenting the input list 
    // based on its value and alternating pattern. It counts contiguous blocks of equal numbers, and then 
    // uses binary search to find the minimal block length that can be formed given the constraint on the 
    // number of such blocks. The second algorithm, on the other hand, checks for an alternating pattern 
    // between even and odd positions of the string using bitwise operations (^), which is not done in the 
    // first algorithm. If either the number of even or odd indexed numbers with alternating values is less 
    // than or equal to k, the solution directly returns 1. The second algorithm is more focused on optimizing 
    // the grouping of the values based on how contiguous segments are formed, while the first algorithm uses 
    // the alternating index check more directly to compute the result. Both algorithms use binary search, 
    // but the logic leading up to the binary search differs.
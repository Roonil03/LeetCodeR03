public class Solution {
    public int LongestValidParentheses(string s) {
        int maxLen = 0;
        int left = 0, right = 0;
        for (int i = 0; i < s.Length; i++) {
            if (s[i] == '(') left++;
            else right++;
            if (left == right) maxLen = Math.Max(maxLen, 2 * right);
            else if (right > left) left = right = 0;
        }
        left = right = 0;
        for (int i = s.Length - 1; i >= 0; i--) {
            if (s[i] == '(') left++;
            else right++;
            if (left == right) maxLen = Math.Max(maxLen, 2 * left);
            else if (left > right) left = right = 0;
        }
        return maxLen;
    }
}

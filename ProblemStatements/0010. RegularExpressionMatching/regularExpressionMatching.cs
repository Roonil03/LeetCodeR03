using System;

public class Solution {
    public bool IsMatch(string s, string p) {
        return IsMatchHelper(s, p, 0, 0);
    }
   
    private bool IsMatchHelper(string s, string p, int i, int j) {
        // If we have reached the end of both strings, we have a match
        if (j == p.Length) {
            return i == s.Length;
        }
       
        // Check if the current position in `s` matches the current position in `p`
        bool match = (i < s.Length && (s[i] == p[j] || p[j] == '.'));
       
        // Handle the '*' character
        if (j + 1 < p.Length && p[j + 1] == '*') {
            // Two possibilities:
            // 1. We skip the '*' and the preceding element
            // 2. We use the '*' to match one or more of the preceding element (if there's a match)
            return IsMatchHelper(s, p, i, j + 2) || (match && IsMatchHelper(s, p, i + 1, j));
        } else {
            // If there's no '*', just move to the next character in both strings
            return match && IsMatchHelper(s, p, i + 1, j + 1);
        }
    }
}
/*
public class Program {
    public static void Main(string[] args) {
        Solution solution = new Solution();

        // Test cases
        TestCase(solution, "aa", "a*", true);        // Should match
        TestCase(solution, "aa", "a", false);       // Should not match
        TestCase(solution, "ab", ".*", true);       // Should match
        TestCase(solution, "aab", "c*a*b", true);   // Should match
        TestCase(solution, "mississippi", "mis*is*p*.", false); // Should not match
        TestCase(solution, "abc", ".*b", true);     // Should match
        TestCase(solution, "abc", ".*c", true);     // Should match
        TestCase(solution, "abc", "a.c", true);     // Should match
        TestCase(solution, "abc", ".*d", false);    // Should not match
        TestCase(solution, "", ".*", true);          // Should match empty string
        TestCase(solution, "", ".*a", false);        // Should not match empty string with 'a'

        // Edge cases
        TestCase(solution, "", "", true);             // Both empty strings should match
        TestCase(solution, "abc", ".*", true);       // Should match
    }

    public static void TestCase(Solution solution, string s, string p, bool expected) {
        bool result = solution.IsMatch(s, p);
        Console.WriteLine($"IsMatch(\"{s}\", \"{p}\") = {result}. Expected: {expected}. Test " + (result == expected ? "Passed" : "Failed"));
    }
}
*/

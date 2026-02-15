# 3844. Longest Almost-Palindromic Substring
## Question Level: Medium
### Description:
You are given a string s consisting of lowercase English letters.

A substring is almost-palindromic if it becomes a palindrome after removing exactly one character from it.

Return an integer denoting the length of the longest almost-palindromic substring in s.

A substring is a contiguous non-empty sequence of characters within a string.

A palindrome is a non-empty string that reads the same forward and backward.

### Examples:
#### Example 1:

Input: s = "`abca`"

Output: 4

Explanation:

Choose the substring "`abca`".
- Remove "`abca`".
- The string becomes "`aba`", which is a palindrome.
- Therefore, "`abca`" is almost-palindromic.
#### Example 2:

Input: s = "`abba`"

Output: 4

Explanation:

Choose the substring "`abba`".
- Remove "`abba`".
- The string becomes "`aba`", which is a palindrome.
- Therefore, "`abba`" is almost-palindromic.
#### Example 3:

Input: s = "`zzabba`"

Output: 5

Explanation:

Choose the substring "`zzabba`".
- Remove "`zabba`".
- The string becomes "`abba`", which is a palindrome.
- Therefore, "`zabba`" is almost-palindromic.

### Constraints:

- 2 <= `s.length` <= 2500
- `s` consists of only lowercase English letters.

### *Weekly Contest Question*

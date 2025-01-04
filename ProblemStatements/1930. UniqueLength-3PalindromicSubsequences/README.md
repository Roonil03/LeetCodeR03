# 1930. Unique Length-3 Palindromic Subsequences
## Question Level: Medium
### Description:
Given a string s, return the number of unique palindromes of length three that are a subsequence of s.

Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.

A palindrome is a string that reads the same forwards and backwards.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "`ace`" is a subsequence of "<b>a</b>b<b>c</b>d<b>e</b>".

### Examples:
#### Example 1:

Input: s = "`aabca`"
Output: 3
Explanation: The 3 palindromic subsequences of length 3 are:
- "`aba`" (subsequence of "<b>a</b>a<b>b</b>c<b>a</b>")
- "`aaa`" (subsequence of "<b>a</b><b>a</b>bc<b>a</b>")
- "`aca`" (subsequence of "<b>a</b>ab<b>c</b><b>a</b>")
#### Example 2:

Input: s = "`adc`"
Output: 0
Explanation: There are no palindromic subsequences of length 3 in "adc".
#### Example 3:

Input: s = "`bbcbaba`"
Output: 4
Explanation: The 4 palindromic subsequences of length 3 are:
- "`bbb`" (subsequence of "<b>b</b><b>b</b>c<b>b</b>aba")
- "`bcb`" (subsequence of "<b>b</b>b<b>c</b><b>b</b>aba")
- "`bab`" (subsequence of "<b>b</b>bcb<b>a</b><b>b</b>a")
- "`aba`" (subsequence of "bbcb<b>a</b><b>b</b><b>a</b>")

### Constraints:

- 3 <= `s.length` <= 10^5
- s consists of only lowercase English letters.

### <i>Concepts Used:
- Hash Table
- String
- Bit Manipulation
- Prefix Sum</i>

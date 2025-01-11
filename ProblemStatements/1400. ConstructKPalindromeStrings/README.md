# 1400. Construct K Palindrome Strings
## Question Level: Medium
### Description:
Given a string s and an integer k, return true if you can use all the characters in s to construct k palindrome strings or false otherwise.

### Examples:
#### Example 1:

Input: s = "`annabelle`", k = 2<br>
Output: true<br>
Explanation: You can construct two palindromes using all characters in s.<br>
Some possible constructions "`anna`" + "`elble`", "`anbna`" + "`elle`", "`anellena`" + "`b`"<br>
#### Example 2:

Input: s = "`leetcode`", k = 3<br>
Output: false<br>
Explanation: It is impossible to construct 3 palindromes using all the characters of s.<br>
#### Example 3:

Input: s = "`true`", k = 4<br>
Output: true<br>
Explanation: The only possible solution is to put each character in a separate string.<br>

### Constraints:

- 1 <= `s.length` <= 10^5
- s consists of lowercase English letters.
- 1 <= k <= 10^5

### <i>Concepts Used:
- Hash Table
- String
- Greedy
- Counting </i>
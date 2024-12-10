# 2981. Find Longest Special Substring That Occurs Thrice I
## Question Level: Medium
### Description:
You are given a string `s` that consists of lowercase English letters.

A string is called special if it is made up of only a single character. For example, the string `"abc"` is not special, whereas the strings `"ddd"`, `"zz"`, and `"f"` are special.

Return the length of the longest special substring of `s` which occurs at least thrice, or -1 if no special substring occurs at least thrice.

A substring is a contiguous non-empty sequence of characters within a string.

### Examples:
#### Example 1:

Input: s = `"aaaa"`<br>
Output: 2<br>
Explanation: The longest special substring which occurs thrice is `"aa"`: substrings "<b>aa</b>aa", "a<b>aa</b>a", and "aa<b>aa</b>".
It can be shown that the maximum length achievable is 2.<br>
#### Example 2:

Input: s = `"abcdef"`<br>
Output: -1<br>
Explanation: There exists no special substring which occurs at least thrice. Hence return -1.<br>
#### Example 3:

Input: s = `"abcaba"`<br>
Output: 1<br>
Explanation: The longest special substring which occurs thrice is `"a"`: substrings "<b>a</b>bcaba", "abc<b>a</b>ba", and "abcab<b>a</b>".
It can be shown that the maximum length achievable is 1.<br>

### Constraints:

- 3 <= `s.length` <= 50
- s consists of only lowercase English letters.

### <i>Concepts Used:
- Hash Table
- String
- Binary Search
- Sliding Window
- Counting </i>
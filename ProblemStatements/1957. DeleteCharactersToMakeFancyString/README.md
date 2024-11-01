# 1957. Delete Characters to Make Fancy String
## Question Level: Easy
### Description:
A fancy string is a string where no three consecutive characters are equal.

Given a string s, delete the minimum possible number of characters from s to make it fancy.

Return the final string after the deletion. It can be shown that the answer will always be unique.

### Examples:
<b>Example 1:</b><br>
Input: s = ``"leeetcode"``<br>
Output: ``"leetcode"``<br>
Explanation:
- Remove an 'e' from the first group of 'e's to create "leetcode".
- No three consecutive characters are equal, so return "leetcode".

<b>Example 2:</b><br>
Input: s = ``"aaabaaaa"``<br>
Output: ``"aabaa"``<br>
Explanation:
- Remove an 'a' from the first group of 'a's to create "aabaaaa".
- Remove two 'a's from the second group of 'a's to create "aabaa".
- No three consecutive characters are equal, so return "aabaa".

<b>Example 3:</b><br>
Input: s = ``"aab"``<br>
Output: ``"aab"``<br>
Explanation: 
- No three consecutive characters are equal, so return "aab".

### Constraints:

1 <= s.length <= 10^5
s consists only of lowercase English letters.

### <i>Concepts Used:
- String
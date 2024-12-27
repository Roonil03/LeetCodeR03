# 1316. Distinct Echo Substrings
## Question Level: Medium
### Description:
Return the number of distinct non-empty substrings of text that can be written as the concatenation of some string with itself (i.e. it can be written as `a + a` where a is some string).

### Examples:
#### Example 1:

Input: text = `"abcabcabc"`<br>
Output: 3<br>
Explanation: The 3 substrings are `"abcabc"`, `"bcabca"` and `"cabcab"`.<br>
#### Example 2:

Input: text = `"leetcodeleetcode"`<br>
Output: 2<br>
Explanation: The 2 substrings are `"ee"` and `"leetcodeleetcode"`.<br>

### Constraints:

- 1 <= `text.length` <= 2000
- text has only lowercase English letters.

### <i>Concepts Used:
- String
- Trie
- Rolling Hash
- Hash Function </i>
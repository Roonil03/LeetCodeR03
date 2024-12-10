# 44. Wildcard Matching
## Question Level: Hard
### Description:
Given an input string `(s)` and a pattern `(p)`, implement wildcard pattern matching with support for `'?'` and `'*'` where:

- `'?'` Matches any single character.
- `'*'` Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

### Examples:
#### Example 1:

Input: s = `"aa"`, p = `"a"`<br>
Output: false<br>
Explanation: `"a"` does not match the entire string `"aa"`.<br>
#### Example 2:

Input: s = `"aa"`, p = `"*"`<br>
Output: true<br>
Explanation: `'*'` matches any sequence.<br>
#### Example 3:

Input: s = `"cb"`, p = `"?a"`<br>
Output: false<br>
Explanation: `'?'` matches `'c'`, but the second letter is `'a'`, which does not match `'b'`.<br>

### Constraints:

- 0 <= `s.length`, `p.length` <= 2000
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'?'` or `'*'`.

### <i>Concepts Used:
- String
- Dynamic Programming
- Greedy
- Recursion </i>
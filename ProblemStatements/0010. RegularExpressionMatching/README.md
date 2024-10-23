# 10. Regular Expression Matching
## Question Level: Hard
### Description:
Given an input string s and a pattern p, implement regular expression matching with support for ``'.'`` and ``'*'`` where:
- ``'.'`` Matches any single character.
- ``'*'`` Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

### Examples:
<b>Example 1:</b><br>
Input: s = ``"aa"``, p = ``"a"``<br>
Output: false<br>
Explanation: "a" does not match the entire string "aa".<br>

<b>Example 2:</b><br>
Input: s = ``"aa"``, p = ``"a*"``<br>
Output: true<br>
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

<b>Example 3:</b><br>
Input: s = ``"ab"``, p = ``".*"``<br>
Output: true<br>
Explanation: ``".*"`` means "zero or more ``(*)`` of any character ``(.)``".<br>

### Constraints:

- 1 <= s.length <= 20
- 1 <= p.length <= 20
- s contains only lowercase English letters.
- p contains only lowercase English letters, ``'.'``, and ``'*'``.
- It is guaranteed for each appearance of the character ``'*'``, there will be a previous valid character to match.

### <i>Concepts used:
- String
- Dynamic Programming
- Recursion </i>
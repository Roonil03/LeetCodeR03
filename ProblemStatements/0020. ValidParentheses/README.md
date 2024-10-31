# 20. Valid Parentheses
## Question Level: Easy
### Description:
Given a string s containing just the characters ``'('``, ``')'``, ``'{'``, ``'}'``, ``'['`` and ``']'``, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

### Examples:
<b>Example 1:</b><br>
Input: s = "()"<br>
Output: true<br>

<b>Example 2:</b><br>
Input: s = "()[]{}"<br>
Output: true<br>

<b>Example 3:</b><br>
Input: s = "(]"<br>
Output: false<br>

<b>Example 4:</b><br>
Input: s = "([])"<br>
Output: true<br>

### Constraints:
- 1 <= `s.length`` <= 10^4
- s consists of parentheses only ``'()[]{}'``.

### <i>Concepts Used:
- String
- Stack </i>
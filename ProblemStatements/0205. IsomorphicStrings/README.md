# 205. Isomorphic Strings
## Question Level: Easy
### Description:
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

### Examples:
#### Example 1:
Input: s = "`egg`", t = "`add`"<br>
Output: true<br>
Explanation:<br>
The strings s and t can be made identical by:<br>
Mapping '`e`' to '`a`'.<br>
Mapping '`g`' to '`d`'.<br>
#### Example 2:
Input: s = "`foo`", t = "`bar`"<br>
Output: false<br>
Explanation:<br>
The strings s and t can not be made identical as '`o`' needs to be mapped to both '`a`' and '`r`'.<br>
#### Example 3:
Input: s = "`paper`", t = "`title`"<br>
Output: true<br>

### Constraints:

- 1 <= `s.length` <= 5 * 10^4
- `t.length` == `s.length`
- s and t consist of any valid ascii character.

### <i>Concepts Used:
- Hash Table
- String</i>
 
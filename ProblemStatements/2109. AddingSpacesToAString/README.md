# 2109. Adding Spaces to a String
## Question Level: Medium
### Description:
You are given a 0-indexed string s and a 0-indexed integer array spaces that describes the indices in the original string where spaces will be added. Each space should be inserted before the character at the given index.

For example, given s = "`EnjoyYourCoffee`" and spaces = `[5, 9]`, we place spaces before '`Y`' and '`C`', which are at indices 5 and 9 respectively. Thus, we obtain `"Enjoy Your Coffee"`.<br>
Return the modified string after the spaces have been added.

### Examples:
#### Example 1:

Input: s = `"LeetcodeHelpsMeLearn"`, spaces = `[8,13,15]`<br>
Output: `"Leetcode Helps Me Learn"`<br>
Explanation: <br>
The indices 8, 13, and 15 correspond to the underlined characters in "Leetcode<b>H</b>elps<b>M</b>e<b>L</b>earn".<br>
We then place spaces before those characters.<br>
#### Example 2:

Input: s = `"icodeinpython"`, spaces = `[1,5,7,9]`<br>
Output: `"i code in py thon"`<br>
Explanation:<br>
The indices 1, 5, 7, and 9 correspond to the underlined characters in "i<b>c</b>ode<b>i</b>n<b>p</b>y<b>t</b>hon".<br>
We then place spaces before those characters.<br>
#### Example 3:

Input: s = `"spacing"`, spaces = `[0,1,2,3,4,5,6]`<br>
Output: `" s p a c i n g"`<br>
Explanation:<br>
We are also able to place spaces before the first character of the string.<br>

### Constraints:
- 1 <= `s.length` <= 3 * 10^5
- s consists only of lowercase and uppercase English letters.
- 1 <= `spaces.length` <= 3 * 10^5
- 0 <= `spaces[i]` <= `s.length - 1`
- All the values of spaces are strictly increasing.

### <i>Concepts Used:
- Array
- Two Pointers
- String
- Simulation </i>
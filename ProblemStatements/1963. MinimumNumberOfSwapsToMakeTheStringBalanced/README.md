# 1963. Minimum Number of Swaps to Make the String Balanced
## Question Level: Medium
### Description:
You are given a 0-indexed string s of even length n. The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.

A string is called balanced if and only if:
- It is the empty string, or
- It can be written as AB, where both A and B are balanced strings, or
- It can be written as [C], where C is a balanced string.<br>
You may swap the brackets at any two indices any number of times.

Return the minimum number of swaps to make s balanced.

### Examples:
<b>Example 1:</b><br>
Input: s = "][]["<br>
Output: 1<br>
Explanation: You can make the string balanced by swapping index 0 with index 3.<br>
The resulting string is "[[]]".<br>

<b>Example 2:</b><br>
Input: s = "]]][[["<br>
Output: 2<br>
Explanation: You can do the following to make the string balanced:<br>
- Swap index 0 with index 4. s = "[]][][".<br>
- Swap index 1 with index 5. s = "[[][]]".<br>
The resulting string is "[[][]]".<br>

<b>Example 3:</b><br>
Input: s = "[]"<br>
Output: 0<br>
Explanation: The string is already balanced.<br>

### Constraints:
- n == s.length
- 2 <= n <= 10^6
- n is even.
- s[i] is either '[' or ']'.
- The number of opening brackets '[' equals n / 2, and the number of closing brackets ']' equals n / 2.
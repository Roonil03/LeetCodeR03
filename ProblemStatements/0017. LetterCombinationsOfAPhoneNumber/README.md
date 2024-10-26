# 17. Letter Combinations of a Phone Number
## Question Level: Medium
### Description:
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
<img src = "https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png"><br>

### Examples:
<b>Example 1:</b><br>
Input: digits = ``"23"``<br>
Output: ``["ad","ae","af","bd","be","bf","cd","ce","cf"]``<br>

<b>Example 2:</b><br>
Input: digits = ``""``<br>
Output: ``[]``<br>

<b>Example 3:</b><br>
Input: digits = ``"2"``<br>
Output: ``["a","b","c"]``<br>

### Constraints:
- 0 <= digits.length <= 4
- digits[i] is a digit in the range ``['2', '9']``.

### <i>Concepts Used:
- Hash Table
- String
- Backtracking </i>
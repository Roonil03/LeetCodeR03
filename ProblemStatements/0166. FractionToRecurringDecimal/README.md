# 166. Fraction to Recurring Decimal
## Question Level: Medium
### Description:
Given two integers representing the `numerator` and `denominator` of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return any of them.

It is guaranteed that the length of the answer string is less than 10<sup>4</sup> for all the given inputs.

### Examples:
#### Example 1:

Input: numerator = 1, denominator = 2  
Output: "`0.5`"
#### Example 2:

Input: numerator = 2, denominator = 1  
Output: "`2`"
#### Example 3:

Input: numerator = 4, denominator = 333  
Output: "`0.(012)`"

### Constraints:

- -2<sup>31</sup> <= `numerator`, `denominator` <= 2<sup>31</sup> - 1
- denominator != 0

### <i>Concepts Used:
- Hash Table
- Math
- String</i>
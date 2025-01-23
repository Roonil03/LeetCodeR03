# 89. Gray Code
## Question Level: Medium
### Description:
An n-bit gray code sequence is a sequence of 2^n integers where:

- Every integer is in the inclusive range `[0, 2^n - 1]`,
- The first integer is 0,
- An integer appears no more than once in the sequence,
- The binary representation of every pair of adjacent integers differs by exactly one bit, and
- The binary representation of the first and last integers differs by exactly one bit.

Given an integer n, return any valid n-bit gray code sequence.

### Examples:
#### Example 1:

Input: n = 2<br>
Output: `[0,1,3,2]`<br>
Explanation:<br>
The binary representation of `[0,1,3,2]` is `[00,01,11,10]`.
- 0<b>0</b> and 0<b>1</b> differ by one bit
- <b>0</b>1 and <b>1</b>1 differ by one bit
- 1<b>1</b> and 1<b>0</b> differ by one bit
- <b>1</b>0 and <b>0</b>0 differ by one bit

`[0,2,3,1]` is also a valid gray code sequence, whose binary representation is `[00,10,11,01]`.
- <b>0</b>0 and <b>1</b>0 differ by one bit
- 1<b>0</b> and 1<b>1</b> differ by one bit
- <b>1</b>1 and <b>0</b>1 differ by one bit
- 0<b>1</b> and 0<b>0</b> differ by one bit
#### Example 2:

Input: n = 1<br>
Output: `[0,1]`<br>

### Constraints:

1 <= n <= 16

### <i>Concepts Used:
- Math
- Backtracking
- Bit Manipulation </i>
 

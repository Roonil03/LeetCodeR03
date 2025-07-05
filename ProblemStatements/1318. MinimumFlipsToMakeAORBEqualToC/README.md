# 1318. Minimum Flips to Make a OR b Equal to c
## Question Level: Medium
### Description:
Given 3 positives numbers `a`, `b` and `c`. Return the minimum flips required in some bits of a and b to make (` a OR b == c `). (bitwise OR operation).  
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/01/06/sample_3_1676.png"><br>
Input: a = 2, b = 6, c = 5  
Output: 3  
Explanation: After flips a = 1 , b = 4 , c = 5 such that (`a OR b == c`)
#### Example 2:

Input: a = 4, b = 2, c = 7  
Output: 1  
#### Example 3:

Input: a = 1, b = 2, c = 3  
Output: 0  

### Constraints:

- 1 <= a <= 10<sup>9</sup>
- 1 <= b <= 10<sup>9</sup>
- 1 <= c <= 10<sup>9</sup>


### <i>Concepts Used:
- Bit Manipulation</i>
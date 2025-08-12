# 2787. Ways to Express an Integer as Sum of Powers
## Question Level: Medium
### Description:
Given two positive integers `n` and `x`.

Return the number of ways `n` can be expressed as the sum of the x<sup>th</sup> power of unique positive integers, in other words, the number of sets of unique integers `[n<`sub>`1`</sub>`, n`<sub>`2`</sub>`, ..., n`<sub>`k`</sub>`]` where n = `n`<sub>`1`</sub><sup>`x`</sup> + `n`<sub>`2`</sub><sup>`x`</sup> + ... + `n`<sub>`k`</sub><sup>`x`</sup>.

Since the result can be very large, return it modulo 10<sup>9</sup> + 7.

For example, if `n = 160` and `x = 3`, one way to express n is n = `2`<sup>`3`</sup>` + 3`<sup>`3`</sup>` + 5`<sup>`3`</sup>.

### Examples:
#### Example 1:

Input: n = `10`, x = `2`  
Output: 1  
Explanation: We can express n as the following: n = 3<sup>2</sup> + 1<sup>2</sup> = 10.  
It can be shown that it is the only way to express 10 as the sum of the 2nd power of unique integers.  
#### Example 2:

Input: n = 4, x = 1  
Output: 2  
Explanation: We can express n in the following ways:  
- n = 4<sup>1</sup> = 4.
- n = 3<sup>1</sup> + 1<sup>1</sup> = 4.

### Constraints:

- 1 <= `n` <= 300
- 1 <= `x` <= 5
### <i>Concepts Used:
- Dynamic Programming</i>
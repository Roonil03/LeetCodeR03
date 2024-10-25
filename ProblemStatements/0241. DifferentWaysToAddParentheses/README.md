# 241. Different Ways to Add Parentheses
## Question Level: Medium
### Description:
Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.
The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 10^4.

### Examples:
Example 1:<br>
Input: expression = "2-1-1"<br>
Output: [0,2]<br>
Explanation:<br>
((2-1)-1) = 0 <br>
(2-(1-1)) = 2<br>

Example 2:<br>
Input: expression = "2*3-4*5"<br>
Output: [-34,-14,-10,-10,10]<br>
Explanation:<br>
(2*(3-(4*5))) = -34 <br>
((2*3)-(4*5)) = -14 <br>
((2*(3-4))*5) = -10 <br>
(2*((3-4)*5)) = -10 <br>
(((2*3)-4)*5) = 10<br>

### Constraints:
- 1 <= expression.length <= 20
- expression consists of digits and the operator '+', '-', and '*'.
- All the integer values in the input expression are in the range [0, 99].
- The integer values in the input expression do not have a leading '-' or '+' denoting the sign.
# 1444. Number of Ways of Cutting a Pizza
## Question Level: Hard
### Description:
Given a rectangular pizza represented as a rows x cols matrix containing the following characters: '`A`' (an apple) and '`.`' (empty cell) and given the integer k. You have to cut the pizza into k pieces using k-1 cuts. 

For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces. If you cut the pizza vertically, give the left part of the pizza to a person. If you cut the pizza horizontally, give the upper part of the pizza to a person. Give the last piece of pizza to the last person.

Return the number of ways of cutting the pizza such that each piece contains at least one apple. Since the answer can be a huge number, return this modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/04/23/ways_to_cut_apple_1.png"><br>
Input: pizza = `["A..","AAA","..."]`, k = 3<br>
Output: 3 <br>
Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.<br>
#### Example 2:

Input: pizza = `["A..","AA.","..."]`, k = 3<br>
Output: 1<br>
#### Example 3:

Input: pizza = `["A..","A..","..."]`, k = 1<br>
Output: 1<br>

### Constraints:

- 1 <= `rows`, `cols` <= 50
- `rows` == `pizza.length`
- `cols` == `pizza[i].length`
- 1 <= k <= 10
- `pizza` consists of characters '`A`' and '`.`' only.

### <i>Concepts Used:
- Array
- Dynamic Programming
- Memoization
- Matrix
- Prefix Sum </i>
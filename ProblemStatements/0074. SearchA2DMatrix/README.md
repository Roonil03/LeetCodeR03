# 74. Search a 2D Matrix
## Question Level: Medium
### Description:
You are given an m x n integer matrix matrix with the following two properties:
- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in `O(log(m * n))` time complexity.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/10/05/mat.jpg"><br>
Input: matrix = `[[1,3,5,7],[10,11,16,20],[23,30,34,60]]`, target = 3<br>
Output: true<br>
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg"><br>
Input: matrix = `[[1,3,5,7],[10,11,16,20],[23,30,34,60]]`, target = 13<br>
Output: false<br>


### Constraints:

- m == `matrix.length`
- n == `matrix[i].length`
- 1 <= `m`, `n` <= 100
- -10^4 <= `matrix[i][j]`, `target` <= 10^4

### <i>Concepts Used:
- Array
- Binary Search
- Matrix </i>
# Biweekly Contest 150
## Ranking: 
## Question 1:
<i>Sum of Good Numbers</i>
### Difficulty: Easy
### Points: 3
### Description:
Given an array of integers nums and an integer k, an element `nums[i]` is considered good if it is strictly greater than the elements at indices `i - k` and `i + k` (if those indices exist). If neither of these indices exists, `nums[i]` is still considered good.

Return the sum of all the good elements in the array.

### Testcases:
#### Example 1:
Input: nums = `[1,3,2,1,5,4]`, k = 2  
Output: 12  
Explanation:  
The good numbers are `nums[1]` = 3, `nums[4]` = 5, and `nums[5]` = 4 because they are strictly greater than the numbers at indices `i - k` and `i + k`.

#### Example 2:
Input: nums = `[2,1]`, k = 1  
Output: 2  
Explanation:  
The only good number is `nums[0]` = 2 because it is strictly greater than `nums[1]`.

### Constraints:

2 <= `nums.length` <= 100
1 <= `nums[i]` <= 1000
1 <= k <= `floor(nums.length / 2)`

## Question 2:
<i>Separate Squares I</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given a 2D integer array squares. Each `squares[i]` = `[xi, yi, li]` represents the coordinates of the bottom-left point and the side length of a square parallel to the x-axis.

Find the minimum y-coordinate value of a horizontal line such that the total area of the squares above the line equals the total area of the squares below the line.

Answers within 10-5 of the actual answer will be accepted.

Note: Squares may overlap. Overlapping areas should be counted multiple times.

### Testcases:
#### Example 1:

Input: squares = `[[0,0,1],[2,2,1]]`<br>
Output: 1.00000<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/06/4062example1drawio.png"><br>
Any horizontal line between y = 1 and y = 2 will have 1 square unit above it and 1 square unit below it. The lowest option is 1.

#### Example 2:
Input: squares = `[[0,0,2],[1,1,1]]`<br>
Output: 1.16667<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/15/4062example2drawio.png"><br>
The areas are:
- Below the line: `7/6 * 2` (Red) + `1/6` (Blue) = `15/6` = `2.5`.
- Above the line: `5/6 * 2` (Red) + `5/6` (Blue) = `15/6` = `2.5`.<br>
Since the areas above and below the line are equal, the output is `7/6` = `1.16667`.

### Constraints:

- 1 <= `squares.length` <= 5 * 10^4
- `squares[i]` = `[xi, yi, li]`
- `squares[i].length` == 3
- 0 <= `xi`, `yi` <= 10^9
- 1 <= `li` <= 10^9

## Question 3:
<i>Separate Squares II</i>
### Difficulty: Hard
### Points: 6
### Description:
You are given a 2D integer array squares. Each `squares[i]` = `[xi, yi, li]` represents the coordinates of the bottom-left point and the side length of a square parallel to the x-axis.

Find the minimum y-coordinate value of a horizontal line such that the total area covered by squares above the line equals the total area covered by squares below the line.

Answers within 10^-5 of the actual answer will be accepted.

Note: Squares may overlap. Overlapping areas should be counted only once in this version.

### Testcases:
#### Example 1:

Input: squares = `[[0,0,1],[2,2,1]]`<br>
Output: 1.00000<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/15/4065example1drawio.png"><br>
Any horizontal line between y = 1 and y = 2 results in an equal split, with 1 square unit above and 1 square unit below. The minimum y-value is 1.

#### Example 2:

Input: squares = `[[0,0,2],[1,1,1]]`<br>
Output: 1.00000<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/15/4065example2drawio.png"><br>
Since the blue square overlaps with the red square, it will not be counted again. Thus, the line y = 1 splits the squares into two equal parts.

### Constraints:

- 1 <= `squares.length` <= 5 * 10^4
- `squares[i]` = `[xi, yi, li]`
- `squares[i].length` == 3
- 0 <= `xi`, `yi` <= 10^9
- 1 <= `li` <= 10^9

## Question 4:
<i>Shortest Matching Substring</i>
### Difficulty: Hard
### Points: 7
### Description:
You are given a string s and a pattern string p, where p contains exactly two `'*'` characters.

The `'*'` in p matches any sequence of zero or more characters.

Return the length of the shortest substring in s that matches p. If there is no such substring, return -1.

Note: The empty substring is considered valid.

### Testcases:
#### Example 1:

Input: s = "`abaacbaecebce`", p = "`ba*c*ce`"
Output: 8
Explanation:
The shortest matching substring of p in s is "`baecebce`".

#### Example 2:

Input: s = "`baccbaadbc`", p = "`cc*baa*adb`"
Output: -1
Explanation:
There is no matching substring in s.

#### Example 3:

Input: s = "`a`", p = "`**`"
Output: 0
Explanation:
The empty substring is the shortest matching substring.

#### Example 4:
Input: s = "`madlogic`", p = "`*adlogi*`"
Output: 6
Explanation:
The shortest matching substring of p in s is "`adlogi`".

### Constraints:

- 1 <= `s.length` <= 10^5
- 2 <= `p.length` <= 10^5
- s contains only lowercase English letters.
- p contains only lowercase English letters and exactly two '*'.
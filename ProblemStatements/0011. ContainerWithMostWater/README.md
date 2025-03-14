# 11. Container With Most Water
## Question Level: Medium
### Description:
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are ``(i, 0)`` and ``(i, height[i])``.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

<i>Notice that you may not slant the container.</i>

### Examples:
<b>Example 1:</b><br>
<img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg"><br>

Input: height = [1,8,6,2,5,4,8,3,7]<br>
Output: 49<br>
Explanation: The above vertical lines are represented by array ``[1,8,6,2,5,4,8,3,7]``. In this case, the max area of water (blue section) the container can contain is 49.<br>

<b>Example 2:</b><br>
Input: height = [1,1]<br>
Output: 1<br>

### Constraints:
- n == height.length
- 2 <= n <= 10^5
- 0 <= height[i] <= 10^4

### <i> Concepts used:
- Array
- Two Pointers
- Greedy</i>
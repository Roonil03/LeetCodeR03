# Weekly Contest 427

## Ranking: 4737 / 22325

## Question 1:
<i>Transformed Array</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array nums that represents a circular array. Your task is to create a new array result of the same size, following these rules:

For each index i (where 0 <= i < `nums.length`), perform the following independent actions:
- If `nums[i]` > 0: Start at index i and move `nums[i]` steps to the right in the circular array. Set `result[i]` to the value of the index where you land.
- If `nums[i]` < 0: Start at index i and move `abs(nums[i])` steps to the left in the circular array. Set `result[i]` to the value of the index where you land.
- If `nums[i]` == 0: Set `result[i]` to `nums[i]`.

Return the new array result.

<i>Note: Since nums is circular, moving past the last element wraps around to the beginning, and moving before the first element wraps back to the end.</i>

### Test Cases:
#### Example 1:

Input: nums = `[3,-2,1,1]`<br>
Output: `[1,1,1,3]`<br>
Explanation:
- For `nums[0]` that is equal to 3, If we move 3 steps to right, we reach `nums[3]`. So `result[0]` should be 1.
- For `nums[1]` that is equal to -2, If we move 2 steps to left, we reach `nums[3]`. So `result[1]` should be 1.
- For `nums[2]` that is equal to 1, If we move 1 step to right, we reach `nums[3]`. So `result[2]` should be 1.
- For `nums[3]` that is equal to 1, If we move 1 step to right, we reach `nums[0]`. So `result[3]` should be 3.
#### Example 2:

Input: nums = `[-1,4,-1]`<br>
Output: `[-1,-1,4]`<br>
Explanation:
- For `nums[0]` that is equal to -1, If we move 1 step to left, we reach `nums[2]`. So `result[0]` should be -1.
- For `nums[1]` that is equal to 4, If we move 4 steps to right, we reach `nums[2]`. So `result[1]` should be -1.
- For `nums[2]` that is equal to -1, If we move 1 step to left, we reach `nums[1]`. So `result[2]` should be 4.

### Constraints:

- 1 <= `nums.length` <= 100
- -100 <= `nums[i]` <= 100

## Question 2:
<i>Maximum Area Rectangle With Point Constraints I</i>
### Difficulty: Medium
### Points: 4
### Description:
You are given an array points where `points[i]` = `[xi, yi]` represents the coordinates of a point on an infinite plane.

Your task is to find the maximum area of a rectangle that:
- Can be formed using four of these points as its corners.
- Does not contain any other point inside or on its border.
- Has its edges parallel to the axes.

Return the maximum area that you can obtain or -1 if no such rectangle is possible.

### Test Cases:
#### Example 1:
Input: points = `[[1,1],[1,3],[3,1],[3,3]]`<br>
Output: 4<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png"><br>
We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.

#### Example 2:
Input: points = `[[1,1],[1,3],[3,1],[3,3],[2,2]]`<br>
Output: -1<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png"><br>
There is only one rectangle possible is with points `[1,1]`, `[1,3]`, `[3,1]` and `[3,3]` but `[2,2]` will always lie inside it. Hence, returning -1.

#### Example 3:
Input: points = `[[1,1],[1,3],[3,1],[3,3],[1,2],[3,2]]`<br>
Output: 2<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png"><br>
The maximum area rectangle is formed by the points `[1,3]`, `[1,2]`, `[3,2]`, `[3,3]`, which has an area of 2. Additionally, the points `[1,1]`, `[1,2]`, `[3,1]`, `[3,2]` also form a valid rectangle with the same area.

### Constraints:
- 1 <= `points.length` <= 10
- `points[i].length` == 2
- 0 <= `xi`, `yi` <= 100
- All the given points are unique.


## Question 3:
Maximum Subarray Sum With Length Divisible by K
### Difficulty: Medium
### Points: 5
### Description:
You are given an array of integers nums and an integer k.

Return the maximum sum of a non-empty subarray of nums, such that the size of the subarray is divisible by k.

A subarray is a contiguous non-empty sequence of elements within an array.

### Test Cases:
#### Example 1:
Input: nums = `[1,2]`, k = 1
Output: 3
Explanation:
The subarray `[1, 2]` with sum 3 has length equal to 2 which is divisible by 1.

#### Example 2:
Input: nums = `[-1,-2,-3,-4,-5]`, k = 4
Output: -10
Explanation:
The maximum sum subarray is `[-1, -2, -3, -4]` which has length equal to 4 which is divisible by 4.

#### Example 3:
Input: nums = `[-5,1,2,-3,4]`, k = 2
Output: 4
Explanation:
The maximum sum subarray is `[1, 2, -3, 4]` which has length equal to 4 which is divisible by 2.

### Constraints:
- 1 <= k <= `nums.length` <= 2 * 10^5
- -10^9 <= `nums[i]` <= 10^9


## Question 4:
Maximum Area Rectangle With Point Constraints II
### Difficulty: Hard
### Points: 8
### Description:
There are n points on an infinite plane. You are given two integer arrays xCoord and yCoord where (`xCoord[i]`, `yCoord[i]`) represents the coordinates of the ith point.

Your task is to find the maximum area of a rectangle that:
- Can be formed using four of these points as its corners.
- Does not contain any other point inside or on its border.
- Has its edges parallel to the axes.

Return the maximum area that you can obtain or -1 if no such rectangle is possible.

### Test Cases:
#### Example 1:
Input: xCoord = `[1,1,3,3]`, yCoord = `[1,3,1,3]`<br>
Output: 4<br>
Explanation:
<img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png"><br>
We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.

#### Example 2:
Input: xCoord = `[1,1,3,3,2]`, yCoord = `[1,3,1,3,2]`<br>
Output: -1<br>
Explanation:
<img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png"><br>
There is only one rectangle possible is with points `[1,1]`, `[1,3]`, `[3,1]` and `[3,3]` but `[2,2]` will always lie inside it. <br>Hence, returning -1.

#### Example 3:
Input: xCoord = `[1,1,3,3,1,3]`, yCoord = `[1,3,1,3,2,2]`<br>
Output: 2<br>
Explanation:
<img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png"><br>
The maximum area rectangle is formed by the points `[1,3]`, `[1,2]`, `[3,2]`, `[3,3]`, which has an area of 2. Additionally, the points `[1,1]`, `[1,2]`, `[3,1]`, `[3,2]` also form a valid rectangle with the same area.

### Constraints:

- 1 <= `xCoord.length` == `yCoord.length` <= 2 * 10^5
- 0 <= `xCoord[i]`, `yCoord[i]` <= 8 * 10^7
- All the given points are unique.

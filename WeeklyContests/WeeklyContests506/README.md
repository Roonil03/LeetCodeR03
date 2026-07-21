# Weekly Contest 506
## Ranking: 85 / 36775
## Question 1:
*Check Good Integer*
### Difficulty: Easy
### Points: 3
### Description:
You are given a positive integer n.

Let `digitSum` be the sum of the digits of n, and let `squareSum` be the sum of the squares of the digits of n.

An integer is called good if `squareSum - digitSum >= 50`.

Return true if n is good. Otherwise, return false.

### Examples:
#### Example 1:
Input: n = 1000

Output: false

Explanation:

- The digits of 1000 are 1, 0, 0, and 0.
- The digitSum is 1 + 0 + 0 + 0 = 1.
- The squareSum is 12 + 02 + 02 + 02 = 1.
- The squareSum - digitSum is 1 - 1 = 0. As 0 is not greater than or equal to 50, the output is false.
#### Example 2:

Input: n = 19

Output: true

Explanation:

- The digits of 19 are 1 and 9.
- The digitSum is 1 + 9 = 10.
- The squareSum is 12 + 92 = 1 + 81 = 82.
- The squareSum - digitSum is 82 - 10 = 72. As 72 is greater than or equal to 50, the output is true.

### Constraints:

- 1 <= n <= 10<sup>9</sup>

## Question 2:
*Frequency Balance Subarray*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array ​​​​​​​nums.

Define a frequency balance subarray as follows:

- If the subarray contains only one distinct value, it is frequency balanced.
- Otherwise, there must exist a positive integer f such that every distinct value in the subarray occurs either f or 2 * f times, and both frequencies occur among the distinct values.

Return an integer denoting the length of the longest frequency balance subarray.

### Examples:
#### Example 1:

Input: nums = `[1,2,2,1,2,3,3,3]`

Output: 5

Explanation:

- The longest frequency balance subarray is `[2, 1, 2, 3, 3]`.
- The elements that appear most frequently are 2 and 3, both appearing twice.
- The remaining element 1 appears once, meeting the requirements.
#### Example 2:

Input: nums = `[5,5,5,5]`

Output: 4

Explanation:

- The longest frequency balance subarray is `[5, 5, 5, 5]`.
- The element that appears most frequently is 5.
- There are no other elements meeting the requirements.
#### Example 3:

Input: nums = `[1,2,3,4]`

Output: 1

Explanation:

Since all elements appear only once, the length of the longest frequency balance subarray is 1.

### Constraints:

- 1 <= `nums.length` <= 10<sup>​​​​​​​3</sup>
- 1 <= `nums[i]` <= 10​​​​​​​<sup>9</sup>

## Question 3:
*Maximize Sum of Device Ratings*
### Difficulty: Medium
### Points: 5
### Description:
You are given a 2D integer array units of size m × n where `units[i][j]` represents the capacity of the jth unit in the ith device. Each device contains exactly n units.

The rating of a device is the minimum capacity among all its units.

You may perform the following operation any number of times (including zero):
- Choose a device i that has not been used as a source before.
- Remove exactly one unit from device i and add it to any different device.
- Then mark device i as used, so it cannot be chosen again as a source.

Return the maximum possible sum of the ratings of all devices after any number of such operations.

Note:
- Devices can receive units from multiple devices, regardless of whether they have been selected.
- The rating of an empty device is 0.

### Examples:
#### Example 1:

Input: units = `[[1,3],[2,2]]`

Output: 4

Explanation:
- ​​​​​​​​​​​​​​Select device i = 0 and transfer `units[0][0]` = 1 to device i = 1.
- After the transfer, the ratings are:
    - Device 0 = `[3]`: `rating[0]` = 3
    - Device 1 = `[2, 2, 1]`: `rating[1]` = 1
- Thus, the sum of ratings is 3 + 1 = 4.
#### Example 2:

Input: units = `[[1,2,3],[4,5,6]]`

Output: 6

Explanation:

- Select device i = 1 and transfer `units[1][0]` = 4 to device i = 0.
- After the transfer, the ratings are:
    - Device 0 = `[1, 2, 3, 4]`: `rating[0]` = 1
    - Device 1 = `[5, 6]`: `rating[1]` = 5
- Thus, the sum of ratings is 1 + 5 = 6.
#### Example 3:

Input: units = `[[5,5,5],[1,1,1]]`

Output: 6

Explanation:

No transfers increase the sum of ratings. Thus, the sum of ratings is 5 + 1 = 6.

### Constraints:

- 1 <= `m` == `units.length` <= 10<sup>5</sup>
- 1 <= `n` == `units[i].length` <= 10<sup>5</sup>
- `m * n` <= 2 * 10<sup>5</sup>
- 1 <= `units[i][j]` <= 10<sup>5</sup>

## Question 4:
*Maximum Subarray Sum After at Most K Swaps*
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array `nums` and an integer k.

You are allowed to perform at most k swap operations on the array.

In one swap operation, you may choose any two indices i and j and swap `nums[i]` and `nums[j]`.

Return an integer denoting the maximum possible subarray sum after performing the swaps.

### Examples:
#### Example 1:

Input: nums = `[1,-1,0,2]`, k = 1

Output: 3

Explanation:
- We can swap on indices 1 and 3, resulting in the array `[1, 2, 0, -1]`.
- The subarray `[1, 2]` has a sum of 3, which is the maximum possible subarray sum after at most k = 1​​​​​​​ swap.
#### Example 2:

Input: nums = `[4,3,2,4]`, k = 2

Output: 13

Explanation:

- The maximum possible subarray sum after at most k = 2 swaps is the sum of the entire array, which is 13.

#### Example 3:

Input: nums = `[-1,-2]`, k = 0

Output: -1

Explanation:
- k = 0 swaps are allowed.
- The possible subarrays are `[-1]`, `[-2]`, and `[-1, -2]`, with sums -1, -2, and -3 respectively.
- Among these sums, the maximum is -1.

### Constraints:

- 1 <= `nums.length` <= 1500
- -10<sup>5</sup> <= `nums[i]` <= 10<sup>5</sup>
- 0 <= k <= `nums.length`
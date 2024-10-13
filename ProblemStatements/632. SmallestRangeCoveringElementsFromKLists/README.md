# 632. Smallest Range Covering Elements from K Lists
## Question Level: Hard
### Description:
You have k lists of sorted integers in non-decreasing order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.

### Examples:
<b>Example 1:</b><br>
Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]<br>
Output: [20,24]<br>
Explanation: <br>
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].<br>
List 2: [0, 9, 12, 20], 20 is in range [20,24].<br>
List 3: [5, 18, 22, 30], 22 is in range [20,24].<br>

<b>Example 2:</b><br>
Input: nums = [[1,2,3],[1,2,3],[1,2,3]]<br>
Output: [1,1]<br>

### Constraints:
- nums.length == k
- 1 <= k <= 3500
- 1 <= nums[i].length <= 50
- -10^5 <= nums[i][j] <= 10^5
- nums[i] is sorted in non-decreasing order.
# 1590. Make Sum Divisible by P
## Question Level: Medium
### Description:
Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that the sum of the remaining elements is divisible by p. It is not allowed to remove the whole array.

Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

A subarray is defined as a contiguous block of elements in the array.

### Example:
<b>Example 1:</b><br>
Input: nums = [3,1,4,2], p = 6<br>
Output: 1<br>
Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.<br>

<b>Example 2:</b><br>
Input: nums = [6,3,5,2], p = 9<br>
Output: 2<br>
Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.<br>

<b>Example 3:</b><br>
Input: nums = [1,2,3], p = 3<br>
Output: 0<br>
Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.<br>

### Constraints:
- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9
- 1 <= p <= 10^9
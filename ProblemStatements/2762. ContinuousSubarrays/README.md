# 2762. Continuous Subarrays
## Question Level: Medium
### Description:
You are given a `0`-indexed integer array nums. A subarray of nums is called continuous if:
- Let `i`, `i + 1`, ..., `j` be the indices in the subarray. Then, for each pair of indices `i` <= `i1`, `i2` <= `j`, 0 <= `|nums[i1] - nums[i2]|` <= 2.
Return the total number of continuous subarrays.

A subarray is a contiguous non-empty sequence of elements within an array.

### Examples:
#### Example 1:

Input: nums = `[5,4,2,4]`<br>
Output: 8<br>
Explanation: 
- Continuous subarray of size 1: `[5]`, `[4]`, `[2]`, `[4]`.
- Continuous subarray of size 2: `[5,4]`, `[4,2]`, `[2,4]`.
- Continuous subarray of size 3: `[4,2,4]`.
- Thereare no subarrys of size 4.<br>
Total continuous subarrays = `4 + 3 + 1 = 8.`<br>
It can be shown that there are no more continuous subarrays.
 

#### Example 2:

Input: nums = `[1,2,3]`<br>
Output: 6<br>
Explanation: 
- Continuous subarray of size 1: `[1]`, `[2]`, `[3]`.
- Continuous subarray of size 2: `[1,2]`, `[2,3]`.
- Continuous subarray of size 3: `[1,2,3]`. <br>
Total continuous subarrays = `3 + 2 + 1 = 6`.

### Constraints:

- 1 <= `nums.length` <= 10^5
- 1 <= `nums[i]` <= 10^9

### <i> Concepts Used:
- Sliding Window
- Two Pointers
- Array
- Ordered Map
- Ordered Set
- Monotonic Queue
- Queue
- Heap (Priority Queue)
- Hash Table
- Dynamic Programming
- Binary Search Tree
- Counting
- Segment Tree
- Binary Search
- Tree
- Memoization
- Stack
- Math
- Recursion
- Depth-First Search
- Monotonic Stack
- Greedy
- Iterator </i>
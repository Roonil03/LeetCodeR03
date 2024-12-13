# 2593. Find Score of an Array After Marking All Elements
## Question Level: Medium
### Description:
You are given an array nums consisting of positive integers.

Starting with score = 0, apply the following algorithm:
- Choose the smallest integer of the array that is not marked. If there is a tie, choose the one with the smallest index.
- Add the value of the chosen integer to score.
- Mark the chosen element and its two adjacent elements if they exist.
- Repeat until all the array elements are marked.

Return the score you get after applying the above algorithm.

### Examples:

#### Example 1:

Input: nums = `[2,1,3,4,5,2]`<br>
Output: 7<br>
Explanation: We mark the elements as follows:
- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: `[2,1,3,4,5,2]`.
- 2 is the smallest unmarked element, so we mark it and its left adjacent element: `[2,1,3,4,5,2]`.
- 4 is the only remaining unmarked element, so we mark it: `[2,1,3,4,5,2]`.<br>
Our score is `1 + 2 + 4 = 7`.
#### Example 2:

Input: nums = `[2,3,5,1,3,2]`<br>
Output: 5<br>
Explanation: We mark the elements as follows:
- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: `[2,3,5,1,3,2]`.
- 2 is the smallest unmarked element, since there are two of them, we choose the left-most one, so we mark the one at index 0 and its right adjacent element: `[2,3,5,1,3,2]`.
- 2 is the only remaining unmarked element, so we mark it: `[2,3,5,1,3,2]`.<br>
Our score is `1 + 2 + 2 = 5`.

### Constraints:

- 1 <= `nums.length` <= 10^5
- 1 <= nums[i] <= 10^6

### <i>Concepts Used:
- Heap (Priority Queue)
- Sorting
- Array
- Simulation
- Hash Table
- Ordered Set
- Ordered Map
- Greedy
- Monotonic Stack
- Sliding Window
- Two Pointers
- Stack
- Queue
- Divide and Conquer
- Dynamic Programming
- Doubly-Linked List
- Data Stream
- Radix Sort
- Backtracking
- String
- Tree
- Design
- Hash Function
- Bit Manipulation
- Iterator
- Counting Sort
- Linked List</i>
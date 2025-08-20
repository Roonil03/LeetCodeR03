# 1546. Maximum Number of Non-Overlapping Subarrays With Sum Equals Target
## Question Level: Medium
### Description:
Given an array `nums` and an integer `target`, return the maximum number of non-empty non-overlapping subarrays such that the sum of values in each subarray is equal to `target`.

### Examples:
#### Example 1:

Input: nums = `[1,1,1,1,1]`, target = 2  
Output: 2  
Explanation: There are 2 non-overlapping subarrays `[1,1,1,1,1]` with sum equals to target(2).  
#### Example 2:

Input: nums = `[-1,3,5,1,4,2,-9]`, target = 6  
Output: 2  
Explanation: There are 3 subarrays with sum equal to 6.   
`([5,1], [4,2], [3,5,1,4,2,-9])` but only the first 2 are non-overlapping.  

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- -10<sup>4</sup> <= `nums[i]` <= 10<sup>4</sup>
- 0 <= `target` <= 10<sup>6</sup>

### <i>Concepts Used:
- Array
- Hash Table
- Greedy
- Prefix Sum</i>
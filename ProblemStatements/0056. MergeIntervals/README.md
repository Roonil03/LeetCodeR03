# 56. Merge Intervals
## Question Level: Medium
### Description:
Given an array of intervals where `intervals[i]` = `[starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

### Examples:
#### Example 1:

Input: intervals = `[[1,3],[2,6],[8,10],[15,18]]`<br>
Output: `[[1,6],[8,10],[15,18]]`<br>
Explanation: Since intervals `[1,3]` and `[2,6]` overlap, merge them into `[1,6]`.<br>
#### Example 2:

Input: intervals = `[[1,4],[4,5]]`<br>
Output: `[[1,5]]`<br>
Explanation: Intervals `[1,4]` and `[4,5]` are considered overlapping.<br>

### Constraints:

- 1 <= `intervals.length` <= 10^4
- `intervals[i].length` == 2
- 0 <= `starti` <= `endi` <= 10^4 

### <i>Concepts Used:
- Array
- Sorting </i>
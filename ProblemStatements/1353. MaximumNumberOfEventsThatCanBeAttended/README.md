# 1353. Maximum Number of Events That Can Be Attended
## Question Level: Medium
### Description:
You are given an array of events where `events[i]` = `[startDay`<sub>`i`</sub>`, endDay`<sub>`i`</sub>`]`. Every event i starts at startDay<sub>i</sub> and ends at endDay<sub>i</sub>.

You can attend an event i at any day d where startTime<sub>i</sub> <= d <= endTime<sub>i</sub>. You can only attend one event at any time d.

Return the maximum number of events you can attend.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/02/05/e1.png"><br>
Input: events = `[[1,2],[2,3],[3,4]]`  
Output: 3  
Explanation: You can attend all the three events.  
One way to attend them all is as shown.  
Attend the first event on day 1.  
Attend the second event on day 2.  
Attend the third event on day 3.  
#### Example 2:

Input: events= `[[1,2],[2,3],[3,4],[1,2]]`  
Output: 4  

### Constraints:

- 1 <= `events.length` <= 10<sup>5</sup>
- `events[i].length` == 2
- 1 <= `startDay`<sub>`i`</sub> <= `endDay`<sub>`i`</sub> <= 10<sup>5</sup>

### <i>Concepts Used:
- Array
- Greedy
- Sorting
- Heap (Priority Queue)</i>
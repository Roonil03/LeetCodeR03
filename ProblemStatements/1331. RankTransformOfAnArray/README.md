# 1331. Rank Transform of an Array
## Question Level: Easy
### Description:
Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:
- Rank is an integer starting from 1.
- The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
- Rank should be as small as possible.

### Example:
<b>Example 1:</b><br>
Input: arr = [40,10,20,30]<br>
Output: [4,1,2,3]<br>
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.<br>

<b>Example 2:</b><br>
Input: arr = [100,100,100]<br>
Output: [1,1,1]<br>
Explanation: Same elements share the same rank.<br>

<b>Example 3:</b><br>
Input: arr = [37,12,28,9,100,56,80,5,12]<br>
Output: [5,3,4,2,8,6,7,1,3]<br>


### Constraints:
- 0 <= arr.length <= 10^5
- 10^9 <= arr[i] <= 10^9
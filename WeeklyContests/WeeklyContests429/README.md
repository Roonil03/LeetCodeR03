# Weekly Contests 429
## Rank: 
## Question 1:
<i>Minimum Number of Operations to Make Elements in Array Distinct</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array nums. You need to ensure that the elements in the array are distinct. To achieve this, you can perform the following operation any number of times:
- Remove 3 elements from the beginning of the array. If the array has fewer than 3 elements, remove all remaining elements.

<i>Note that an empty array is considered to have distinct elements.</i>

Return the minimum number of operations needed to make the elements in the array distinct.

### Testcases:
#### Example 1:

Input: nums = `[1,2,3,4,2,3,3,5,7]`<br>
Output: 2<br>
Explanation:
- In the first operation, the first 3 elements are removed, resulting in the array `[4, 2, 3, 3, 5, 7]`.
- In the second operation, the next 3 elements are removed, resulting in the array `[3, 5, 7]`, which has distinct elements.<br>
Therefore, the answer is 2.

#### Example 2:
Input: nums = `[4,5,6,4,4]`<br>
Output: 2<br>
Explanation:
- In the first operation, the first 3 elements are removed, resulting in the array `[4, 4]`.
- In the second operation, all remaining elements are removed, resulting in an empty array.<br>
Therefore, the answer is 2.

#### Example 3:
Input: nums = `[6,7,8,9]`<br>
Output: 0<br>
Explanation:<br>
The array already contains distinct elements. Therefore, the answer is 0.

### Constraints:

- 1 <= `nums.length` <= 100
- 1 <= `nums[i]` <= 100

## Question 2:
<i>Maximum Number of Distinct Elements After Operations</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array `nums` and an integer k.

You are allowed to perform the following operation on each element of the array at most once:
- Add an integer in the range `[-k, k]` to the element.

Return the maximum possible number of distinct elements in nums after performing the operations.

### Testcases:
#### Example 1:
Input: nums = `[1,2,2,3,3,4]`, k = 2<br>
Output: 6<br>
Explanation:<br>
nums changes to `[-1, 0, 1, 2, 3, 4]` after performing operations on the first four elements.

#### Example 2:
Input: nums = `[4,4,4,4]`, k = 1<br>
Output: 3<br>
Explanation:<br>
By adding -1 to `nums[0]` and 1 to `nums[1]`, nums changes to `[3, 5, 4, 4]`.


### Constraints:

- 1 <= `nums.length` <= 10^5
- 1 <= `nums[i]` <= 10^9
- 0 <= `k` <= 10^9

## Question 3:
<i>Smallest Substring With Identical Characters I</i>
### Difficulty: Hard
### Points: 6
### Description:
You are given a binary string s of length n and an integer `numOps`.

You are allowed to perform the following operation on s at most `numOps` times:
- Select any index i (where 0 <= i < n) and flip `s[i]`, i.e., if `s[i]` == `'1'`, change `s[i]` to `'0'` and vice versa.
- You need to minimize the length of the longest substring of s such that all the characters in the substring are identical.

Return the minimum length after the operations.

A substring is a contiguous non-empty sequence of characters within a string.

### Testcases:
#### Example 1:
Input: s = `"000001"`, numOps = 1<br>
Output: 2<br>
Explanation: <br>
By changing `s[2]` to '1', s becomes `"001001"`. The longest substrings with identical characters are `s[0..1]` and `s[3..4]`.

#### Example 2:
Input: s = `"0000"`, numOps = 2<br>
Output: 1<br>
Explanation: <br>
By changing `s[0]` and `s[2]` to '1', s becomes "`1010`".<br>

#### Example 3:
Input: s = `"0101"`, numOps = 0<br>
Output: 1<br>

### Constraints:

- 1 <= `n` == `s.length` <= 1000
- s consists only of '0' and '1'.
- 0 <= `numOps` <= n

## Question 4:
<i>Smallest Substring With Identical Characters II</i>
### Difficulty: Hard
### Points: 7
### Description:
You are given a binary string s of length n and an integer numOps.

You are allowed to perform the following operation on s at most numOps times:
- Select any index i (where 0 <= i < n) and flip `s[i]`. If `s[i]` == '`1`', change `s[i]` to '`0`' and vice versa.

You need to minimize the length of the longest substring of s such that all the characters in the substring are identical.

Return the minimum length after the operations.

A substring is a contiguous non-empty sequence of characters within a string.

### Testcases:
#### Example 1:
Input: s = "`000001`", numOps = 1<br>
Output: 2<br>
Explanation: <br>
By changing `s[2]` to '1', s becomes `"001001"`. The longest substrings with identical characters are `s[0..1]` and `s[3..4]`.<br>

#### Example 2:
Input: s = "`0000`", numOps = 2<br>
Output: 1<br>
Explanation: <br>
By changing `s[0]` and `s[2]` to '`1`', s becomes "`1010`".<br>

#### Example 3:
Input: s = "`0101`", numOps = 0<br>
Output: 1<br>

### Constraints:

- 1 <= `n` == `s.length` <= 10^5
- `s` consists only of '0' and '1'.
- 0 <= `numOps` <= n
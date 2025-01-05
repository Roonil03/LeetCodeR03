# Biweekly Contest 147
## Ranking: 
## Question 1:
<i>Substring Matching Pattern</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given a string s and a pattern string p, where p contains exactly one `'*'` character.

The `'*'` in p can be replaced with any sequence of zero or more characters.

Return true if p can be made a substring of s, and false otherwise.

A substring is a contiguous non-empty sequence of characters within a string.
### Examples:
#### Example 1:
Input: s = "`leetcode`", p = "`ee*e`"
Output: true
Explanation:
By replacing the `'*'` with "tcod", the substring "`eetcode`" matches the pattern.

#### Example 2:
Input: s = "`car`", p = "`c*v`"
Output: false
Explanation:
There is no substring matching the pattern.

#### Example 3:
Input: s = "`luck`", p = "`u*`"
Output: true
Explanation:
The substrings "`u`", "`uc`", and "`uck`" match the pattern.

### Constraints:
- 1 <= `s.length` <= 50
- 1 <= `p.length` <= 50 
- s contains only lowercase English letters.
- p contains only lowercase English letters and exactly one `'*'`

## Question 2:
<i>Design Task Manager</i>
### Difficulty: Medium
### Points: 5
### Description:
There is a task management system that allows users to manage their tasks, each associated with a priority. The system should efficiently handle adding, modifying, executing, and removing tasks.

Implement the TaskManager class:
- `TaskManager(vector<vector<int>>& tasks)` initializes the task manager with a list of user-task-priority triples. Each element in the input list is of the form [userId, taskId, priority], which adds a task to the specified user with the given priority.

- `void add(int userId, int taskId, int priority)` adds a task with the specified taskId and priority to the user with userId. It is guaranteed that taskId does not exist in the system.

- `void edit(int taskId, int newPriority)` updates the priority of the existing taskId to newPriority. It is guaranteed that taskId exists in the system.

- `void rmv(int taskId)` removes the task identified by taskId from the system. It is guaranteed that taskId exists in the system.

- `int execTop()` executes the task with the highest priority across all users. If there are multiple tasks with the same highest priority, execute the one with the highest taskId. After executing, the taskId is removed from the system. Return the userId associated with the executed task. If no tasks are available, return -1.

Note that a user may be assigned multiple tasks.

### Examples:
#### Example 1:

Input:
```
["TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"]
[[[[1, 101, 10], [2, 102, 20], [3, 103, 15]]], [4, 104, 5], [102, 8], [], [101], [5, 105, 15], []]
```
Output:
```
[null, null, null, 3, null, null, 5]
```
Explanation:
```
TaskManager taskManager = new TaskManager([[1, 101, 10], [2, 102, 20], [3, 103, 15]]); // Initializes with three tasks for Users 1, 2, and 3.
taskManager.add(4, 104, 5); // Adds task 104 with priority 5 for User 4.
taskManager.edit(102, 8); // Updates priority of task 102 to 8.
taskManager.execTop(); // return 3. Executes task 103 for User 3.
taskManager.rmv(101); // Removes task 101 from the system.
taskManager.add(5, 105, 15); // Adds task 105 with priority 15 for User 5.
taskManager.execTop(); // return 5. Executes task 105 for User 5.
```

### Constraints:

1 <= `tasks.length` <= 10^5
0 <= `userId` <= 10^5
0 <= `taskId` <= 10^5
0 <= `priority` <= 10^9
0 <= `newPriority` <= 10^9
At most 2 * 10^5 calls will be made in total to `add`, `edit`, `rmv`, and `execTop` methods.


## Question 3:
<i>Longest Subsequence With Decreasing Adjacent Difference</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an array of integers nums.

Your task is to find the length of the longest subsequence seq of nums, such that the absolute differences between consecutive elements form a non-increasing sequence of integers. In other words, for a subsequence `seq0`, `seq1`, `seq2`, ..., `seqm` of nums, `|seq1 - seq0|` >= `|seq2 - seq1|` >= ... >= `|seqm - seqm - 1|`.

Return the length of such a subsequence.

A subsequence is an non-empty array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.


### Examples:
#### Example 1:
Input: nums = `[16,6,3]`<br>
Output: 3<br>
Explanation: <br>
The longest subsequence is `[16, 6, 3]` with the absolute adjacent differences `[10, 3]`.<br>

#### Example 2:
Input: nums = `[6,5,3,4,2,1]`<br>
Output: 4<br>
Explanation:<br>
The longest subsequence is `[6, 4, 2, 1]` with the absolute adjacent differences `[2, 2, 1]`.<br>

#### Example 3:
Input: nums = `[10,20,10,19,10,20]`<br>
Output: 5<br>
Explanation: <br>
The longest subsequence is `[10, 20, 10, 19, 10]` with the absolute adjacent differences `[10, 10, 9, 9]`.<br>

### Constraints:

- 2 <= `nums.length` <= 10^4
- 1 <= `nums[i]` <= 300

## Question 4:
<i>Maximize Subarray Sum After Removing All Occurrences of One Element</i>
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array `nums`.

You can do the following operation on the array at most once:
- Choose any integer x such that `nums` remains non-empty on removing all occurrences of x.
- Remove all occurrences of x from the array.
- Return the maximum subarray sum across all possible resulting arrays.

A subarray is a contiguous non-empty sequence of elements within an array.

### Examples:
#### Example 1:
Input: nums = `[-3,2,-2,-1,3,-2,3]`<br>
Output: 7<br>
Explanation:
- We can have the following arrays after at most one operation:
- The original array is nums = [-3, 2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
- Deleting all occurences of x = -3 results in nums = [2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
- Deleting all occurences of x = -2 results in nums = [-3, 2, -1, 3, 3]. The maximum subarray sum is 2 + (-1) + 3 + 3 = 7.
- Deleting all occurences of x = -1 results in nums = [-3, 2, -2, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
- Deleting all occurences of x = 3 results in nums = [-3, 2, -2, -1, -2]. The maximum subarray sum is 2.<br>
The output is `max(4, 4, 7, 4, 2)` = `7`.

#### Example 2:
Input: nums = `[1,2,3,4]`<br>
Output: 10<br>
Explanation:<br>
It is optimal to not perform any operations.

### Constraints:

- 1 <= `nums.length` <= 10^5
- -10^6 <= `nums[i]` <= 10^6


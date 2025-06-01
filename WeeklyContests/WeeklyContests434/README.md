# Weekly Contests 434
## Ranking: 389 / 34958
## Question 1:
<i>Count Partitions with Even Sum Difference</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array nums of length n.

A partition is defined as an index i where 0 <= i < n - 1, splitting the array into two non-empty subarrays such that:
- Left subarray contains indices `[0, i]`.
- Right subarray contains indices `[i + 1, n - 1]`.

Return the number of partitions where the difference between the sum of the left and right subarrays is even.

### Testcases:
#### Example 1:
Input: nums = `[10,10,3,7,6]`<br>
Output: 4<br>
Explanation:<br>
The 4 partitions are:
- `[10]`, `[10, 3, 7, 6]` with a sum difference of 10 - 26 = -16, which is even.
- `[10, 10]`, `[3, 7, 6]` with a sum difference of 20 - 16 = 4, which is even.
- `[10, 10, 3]`, `[7, 6]` with a sum difference of 23 - 13 = 10, which is even.
- `[10, 10, 3, 7]`, `[6]` with a sum difference of 30 - 6 = 24, which is even.
#### Example 2:
Input: nums = `[1,2,2]`<br>
Output: 0<br>
Explanation:<br>
No partition results in an even sum difference.<br>

#### Example 3:
Input: nums = `[2,4,6,8]`<br>
Output: 3<br>
Explanation:<br>
All partitions result in an even sum difference.<br>


### Constraints:

2 <= n == `nums.length` <= 100
1 <= `nums[i]` <= 100

## Question 2:
<i>Count Mentions Per User</i>
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer numberOfUsers representing the total number of users and an array events of size n x 3.

Each `events[i]` can be either of the following two types:
1. Message Event: ["`MESSAGE`", "`timestampi`", "`mentions_stringi`"]
- This event indicates that a set of users was mentioned in a message at `timestampi`.
- The `mentions_stringi` string can contain one of the following tokens:
    - `id<number>`: where `<number>` is an integer in range `[0,numberOfUsers - 1]`. There can be multiple ids separated by a single whitespace and may contain duplicates. This can mention even the offline users.
    - `ALL`: mentions all users.
    - `HERE`: mentions all online users.
2. Offline Event: ["`OFFLINE`", "`timestampi`", "`idi`"]
- This event indicates that the user idi had become offline at `timestampi` for 60 time units. The user will automatically be online again at time `timestampi + 60`.

Return an array mentions where `mentions[i]` represents the number of mentions the user with id i has across all MESSAGE events.

All users are initially online, and if a user goes offline or comes back online, their status change is processed before handling any message event that occurs at the same timestamp.

Note that a user can be mentioned multiple times in a single message event, and each mention should be counted separately.

### Testcases:
#### Example 1:
Input: numberOfUsers = 2, events = `[["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]`<br>
Output: `[2,2]`<br>
Explanation:<br>
Initially, all users are online.<br>
At timestamp 10, id1 and id0 are mentioned. mentions = `[1,1]`<br>
At timestamp 11, id0 goes offline.<br>
At timestamp 71, id0 comes back online and "HERE" is mentioned. mentions = `[2,2]`<br>

#### Example 2:

Input: numberOfUsers = 2, events = `[["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]`<br>
Output: `[2,2]`<br>
Explanation:<br>
Initially, all users are online.<br>
At timestamp 10, id1 and id0 are mentioned. mentions = `[1,1]`<br>
At timestamp 11, id0 goes offline.<br>
At timestamp 12, "ALL" is mentioned. This includes offline users, so both id0 and id1 are mentioned. mentions = `[2,2]`<br>

#### Example 3:

Input: numberOfUsers = 2, events = `[["OFFLINE","10","0"],["MESSAGE","12","HERE"]]`<br>
Output: `[0,1]`<br>
Explanation:<br>
Initially, all users are online.<br>
At timestamp 10, id0 goes offline.<br>
At timestamp 12, "HERE" is mentioned. Because id0 is still offline, they will not be mentioned. mentions = `[0,1]`<br>

### Constraints:

- 1 <= `numberOfUsers` <= 100
- 1 <= `events.length` <= 100
- `events[i].length` == 3
- `events[i][0]` will be one of MESSAGE or OFFLINE.
- 1 <= `int(events[i][1])` <= 10^5
- The number of `id<number>` mentions in any "`MESSAGE`" event is between 1 and 100.
- 0 <= `<number>` <= numberOfUsers - 1
- It is guaranteed that the user id referenced in the OFFLINE event is online at the time the event occurs.


## Question 3:
<i>Maximum Frequency After Subarray Operation</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an array nums of length n. You are also given an integer k.

You perform the following operation on nums once:
- Select a subarray `nums[i..j]` where 0 <= i <= j <= n - 1.
- Select an integer x and add x to all the elements in `nums[i..j]`.

Find the maximum frequency of the value k after the operation.

A subarray is a contiguous non-empty sequence of elements within an array.

### Testcases:
#### Example 1:
Input: nums = `[1,2,3,4,5,6]`, k = 1<br>
Output: 2<br>
Explanation:<br>
After adding -5 to `nums[2..5]`, 1 has a frequency of 2 in `[1, 2, -2, -1, 0, 1]`.<br>

#### Example 2:
Input: nums = `[10,2,3,4,5,5,4,3,2,2]`, k = 10<br>
Output: 4<br>
Explanation:<br>
After adding 8 to `nums[1..9]`, 10 has a frequency of 4 in `[10, 10, 11, 12, 13, 13, 12, 11, 10, 10]`.<br>

### Constraints:

1 <= n == `nums.length` <= 10^5
1 <= `nums[i]` <= 50
1 <= k <= 50


## Question 4:
<i>Frequencies of Shortest Supersequences</i>
### Difficulty: Hard
### Points: 8
### Description:
You are given an array of strings words. Find all shortest common `supersequences` (SCS) of words that are not permutations of each other.

A shortest common supersequence is a string of minimum length that contains each string in words as a subsequence.

Return a 2D array of integers freqs that represent all the SCSs. Each `freqs[i]` is an array of size 26, representing the frequency of each letter in the lowercase English alphabet for a single SCS. You may return the frequency arrays in any order.

A permutation is a rearrangement of all the characters of a string.

A subsequence is a non-empty string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

### Testcases:
#### Example 1:
Input: words = `["ab","ba"]`<br>
Output: `[[1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]`<br>
Explanation:<br>
The two SCSs are "`aba`" and "`bab`". The output is the letter frequencies for each one.<br>

#### Example 2:
Input: words = `["aa","ac"]`<br>
Output: `[[2,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]`<br>
Explanation:<br>
The two SCSs are "`aac`" and "`aca`". Since they are permutations of each other, keep only "`aac`".<br>

#### Example 3:
Input: words = `["aa","bb","cc"]`<br>
Output: `[[2,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]`<br>
Explanation:<br>
"`aabbcc`" and all its permutations are SCSs.<br>


### Constraints:

- 1 <=` words.length` <= 256
- `words[i].length` == 2
- All strings in words will altogether be composed of no more than 16 unique lowercase letters.
- All strings in words are unique.


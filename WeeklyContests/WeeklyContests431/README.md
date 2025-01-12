# Weekly Contests 431
## Ranking: 8990 / 24190
## Question 1:
<i>Maximum Subarray With Equal Products</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given an array of positive integers `nums`.

An array arr is called product equivalent if `prod(arr)` == `lcm(arr) * gcd(arr)`, where:
- `prod(arr)` is the product of all elements of arr.
- `gcd(arr)` is the GCD of all elements of arr.
- `lcm(arr)` is the LCM of all elements of arr.

Return the length of the longest product equivalent subarray of `nums`.

A subarray is a contiguous non-empty sequence of elements within an array.

The term `gcd(a, b)` denotes the greatest common divisor of a and b.

The term `lcm(a, b)` denotes the least common multiple of a and b.


### Testcases:
#### Example 1:
Input: nums = `[1,2,1,2,1,1,1]`<br>
Output: 5<br>
Explanation: <br>
The longest product equivalent subarray is `[1, 2, 1, 1, 1]`, where `prod([1, 2, 1, 1, 1])` = 2, `gcd([1, 2, 1, 1, 1])` = 1, and `lcm([1, 2, 1, 1, 1])` = 2.

#### Example 2:
Input: nums = `[2,3,4,5,6]`<br>
Output: 3<br>
Explanation: <br>
The longest product equivalent subarray is `[3, 4, 5]`.<br>

#### Example 3:
Input: nums = `[1,2,3,1,4,5,1]`<br>
Output: 5<br>

### Constraints:

- 2 <= `nums.length` <= 100
- 1 <= `nums[i]` <= 10

## Question 2:
<i>Find Mirror Score of a String</i>
### Difficulty: Medium
### Points: 4
### Description:
You are given a string s.

We define the mirror of a letter in the English alphabet as its corresponding letter when the alphabet is reversed. For example, the mirror of '`a`' is '`z`', and the mirror of '`y`' is '`b`'.

Initially, all characters in the string s are unmarked.

You start with a score of 0, and you perform the following process on the string s:
- Iterate through the string from left to right.
- At each index i, find the closest unmarked index j such that j < i and `s[j]` is the mirror of `s[i]`. Then, mark both indices i and j, and add the value `i - j` to the total score.
- If no such index j exists for the index i, move on to the next index without making any changes.

Return the total score at the end of the process.

### Testcases:
#### Example 1:
Input: s = "`aczzx`"<br>
Output: 5<br>
Explanation:
- `i = 0`. There is no index j that satisfies the conditions, so we skip.
- `i = 1`. There is no index j that satisfies the conditions, so we skip.
- `i = 2`. The closest index j that satisfies the conditions is j = 0, so we mark both indices 0 and 2, and then add 2 - 0 = 2 to the score.
- `i = 3`. There is no index j that satisfies the conditions, so we skip.
- `i = 4`. The closest index j that satisfies the conditions is j = 1, so we mark both indices 1 and 4, and then add 4 - 1 = 3 to the score.
#### Example 2:
Input: s = "`abcdef`"<br>
Output: 0<br>
Explanation:<br>
For each index i, there is no index j that satisfies the conditions.

### Constraints:

- 1 <= `s.length` <= 10^5
- s consists only of lowercase English letters.

## Question 3:
<i>Maximum Coins From K Consecutive Bags</i>
### Difficulty: Medium
### Points: 5
### Description:
There are an infinite amount of bags on a number line, one bag for each coordinate. Some of these bags contain coins.

You are given a 2D array coins, where `coins[i]` = `[li, ri, ci]` denotes that every bag from `li` to `ri` contains `ci` coins.

The segments that coins contain are non-overlapping.

You are also given an integer k.

Return the maximum amount of coins you can obtain by collecting k consecutive bags.

### Testcases:
#### Example 1:
Input: coins = `[[8,10,1],[1,3,2],[5,6,4]]`, k = 4<br>
Output: 10<br>
Explanation:<br>
Selecting bags at positions `[3, 4, 5, 6]` gives the maximum number of coins: `2 + 0 + 4 + 4` = `10`.

#### Example 2:
Input: coins = `[[1,10,3]]`, k = 2<br>
Output: 6<br>
Explanation:<br>
Selecting bags at positions `[1, 2]` gives the maximum number of coins: `3 + 3` = `6`.

### Constraints:

- 1 <= `coins.length` <= 10^5
- 1 <= k <= 10^9
- `coins[i]` == `[li, ri, ci]`
- 1 <= `li` <= `ri` <= 10^9
- 1 <= `ci` <= 1000
- The given segments are non-overlapping.

### Question 4:
<i>Maximum Score of Non-overlapping Intervals</i>
### Difficulty: Hard
### Points: 6
### Description:
You are given a 2D integer array intervals, where `intervals[i]` = `[li, ri, weighti]`. Interval i starts at position `li` and ends at `ri`, and has a weight of `weighti`. You can choose up to 4 non-overlapping intervals. The score of the chosen intervals is defined as the total sum of their weights.

Return the lexicographically smallest array of at most 4 indices from intervals with maximum score, representing your choice of non-overlapping intervals.

Two intervals are said to be non-overlapping if they do not share any points. In particular, intervals sharing a left or right boundary are considered overlapping.

An array `a` is lexicographically smaller than an array `b` if in the first position where `a` and `b` differ, array `a` has an element that is less than the corresponding element in `b`.
If the first `min(a.length, b.length)` elements do not differ, then the shorter array is the lexicographically smaller one.

### Testcases:
#### Example 1:
Input: intervals = `[[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]`<br>
Output: `[2,3]`<br>
Explanation:<br>
You can choose the intervals with indices 2, and 3 with respective weights of 5, and 3.

#### Example 2:
Input: intervals = `[[5,8,1],[6,7,7],[4,7,3],[9,10,6],[7,8,2],[11,14,3],[3,5,5]]`<br>
Output: `[1,3,5,6]`<br>
Explanation:<br>
You can choose the intervals with indices 1, 3, 5, and 6 with respective weights of 7, 6, 3, and 5.

### Constraints:

- 1 <= `intevals.length` <= 5 * 10^4
- `intervals[i].length` == 3
- `intervals[i]` = `[li, ri, weighti]`
- 1 <= `li` <= `ri` <= 10^9
- 1 <= `weighti` <= 10^9
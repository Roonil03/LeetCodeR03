# Weekly Contest 489
## Ranking:
## Question 1:
*Toggle Light Bulbs*
### Difficulty: Easy
### Points: 3
### Description:
You are given an array bulbs of integers between 1 and 100.

There are 100 light bulbs numbered from 1 to 100. All of them are switched off initially.

For each element `bulbs[i]` in the array bulbs:
- If the `bulbs[i]`<sup>th</sup> light bulb is currently off, switch it on.
- Otherwise, switch it off.

Return the list of integers denoting the light bulbs that are on in the end, sorted in ascending order. If no bulb is on, return an empty list.

### Examples:
#### Example 1:

Input: bulbs = `[10,30,20,10]`

Output: `[20,30]`

Explanation:

- The `bulbs[0] = 10th` light bulb is currently off. We switch it on.
- The `bulbs[1] = 30th` light bulb is currently off. We switch it on.
- The `bulbs[2] = 20th` light bulb is currently off. We switch it on.
- The `bulbs[3] = 10th` light bulb is currently on. We switch it off.
- In the end, the `20th` and the `30th` light bulbs are on.
#### Example 2:

Input: bulbs = `[100,100]`

Output: `[]`

Explanation:
- The `bulbs[0] = 100th` light bulb is currently off. We switch it on.
- The `bulbs[1] = 100th` light bulb is currently on. We switch it off.
- In the end, no light bulb is on.


### Constraints:

1 <= `bulbs.length` <= 100
1 <= `bulbs[i]` <= 100

## Question 2:
*First Element with Unique Frequency*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums.

Return an integer denoting the first element (scanning from left to right) in nums whose frequency is unique. That is, no other integer appears the same number of times in nums. If there is no such element, return -1.

### Examples:
#### Example 1:

Input: nums = `[20,10,30,30]`

Output: 30

Explanation:
- 20 appears once.
- 10 appears once.
- 30 appears twice.
- The frequency of 30 is unique because no other integer appears exactly twice.
#### Example 2:

Input: nums = `[20,20,10,30,30,30]`

Output: 20

Explanation:
- 20 appears twice.
- 10 appears once.
- 30 appears 3 times.
- The frequency of 20, 10, and 30 are unique. The first element that has unique frequency is 20.
#### Example 3:

Input: nums = `[10,10,20,20]`

Output: -1

Explanation:
- 10 appears twice.
- 20 appears twice.
- No element has a unique frequency.


### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>5</sup>

## Question 3:
*Longest Almost-Palindromic Substring*
### Difficulty: Medium
### Points: 5
### Description:
You are given a string s consisting of lowercase English letters.

A substring is almost-palindromic if it becomes a palindrome after removing exactly one character from it.

Return an integer denoting the length of the longest almost-palindromic substring in s.

A substring is a contiguous non-empty sequence of characters within a string.

A palindrome is a non-empty string that reads the same forward and backward.

### Examples:
#### Example 1:

Input: s = "`abca`"

Output: 4

Explanation:

Choose the substring "`abca`".
- Remove "`abca`".
- The string becomes "`aba`", which is a palindrome.
- Therefore, "`abca`" is almost-palindromic.
#### Example 2:

Input: s = "`abba`"

Output: 4

Explanation:

Choose the substring "`abba`".
- Remove "`abba`".
- The string becomes "`aba`", which is a palindrome.
- Therefore, "`abba`" is almost-palindromic.
#### Example 3:

Input: s = "`zzabba`"

Output: 5

Explanation:

Choose the substring "`zzabba`".
- Remove "`zabba`".
- The string becomes "`abba`", which is a palindrome.
- Therefore, "`zabba`" is almost-palindromic.

### Constraints:

- 2 <= `s.length` <= 2500
- `s` consists of only lowercase English letters.

## Question 4:
*Maximum Subarray XOR with Bounded Range*
### Difficulty: Hard
### Points: 6
### Description:
You are given a non-negative integer array nums and an integer k.

You must select a subarray of nums such that the difference between its maximum and minimum elements is at most k. The value of this subarray is the bitwise XOR of all elements in the subarray.

Return an integer denoting the maximum possible value of the selected subarray.

A subarray is a contiguous non-empty sequence of elements within an array.

### Examples:
#### Example 1:

Input: nums = `[5,4,5,6]`, k = 2

Output: 7

Explanation:
- Select the subarray [5, **4, 5, 6**].
- The difference between its maximum and minimum elements is 6 - 4 = 2 <= k.
- The value is 4 XOR 5 XOR 6 = 7.
#### Example 2:

Input: nums = `[5,4,5,6]`, k = 1

Output: 6

Explanation:
- Select the subarray [5, 4, 5, **6**].
- The difference between its maximum and minimum elements is 6 - 6 = 0 <= k.
- The value is 6.

### Constraints:

- 1 <= `nums.length` <= 4 * 10<sup>4</sup>
- 0 <= `nums[i]` < 2<sup>15</sup>
- 0 <= `k` < 2<sup>15</sup>
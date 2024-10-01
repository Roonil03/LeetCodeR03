# 1497. Check If Array Pairs Are Divisible by k
## Question Level: Medium
### Description:
Given an array of integers arr of even length n and an integer k.

We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

Return true If you can find a way to do that or false otherwise.

### Examples:
<b>Example 1:</b><br>
Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5<br>
Output: true<br>
Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).<br>

<b>Example 2:</b><br>
Input: arr = [1,2,3,4,5,6], k = 7<br>
Output: true<br>
Explanation: Pairs are (1,6),(2,5) and(3,4).<br>

<b>Example 3:</b><br>
Input: arr = [1,2,3,4,5,6], k = 10<br>
Output: false<br>
Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.<br>

### Constraints:

- arr.length == n
- 1 <= n <= 105
- n is even.
- -109 <= arr[i] <= 109
- 1 <= k <= 105
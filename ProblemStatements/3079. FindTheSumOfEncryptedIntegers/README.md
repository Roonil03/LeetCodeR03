# 3079. Find the Sum of Encrypted Integers
## Question Level: Easy
### Description:
You are given an integer array nums containing positive integers. We define a function encrypt such that `encrypt(x)` replaces every digit in x with the largest digit in x. For example, `encrypt(523)` = `555` and `encrypt(213)` = `333`.

Return the sum of encrypted elements.

### Examples:
#### Example 1:
Input: nums = `[1,2,3]`<br>
Output: 6<br>
Explanation: The encrypted elements are `[1,2,3]`. The sum of encrypted elements is `1` + `2` + `3` == `6`.<br>

#### Example 2:
Input: nums = `[10,21,31]`<br>
Output: 66<br>
Explanation: The encrypted elements are `[11,22,33]`. The sum of encrypted elements is `11` + `22` + `33` == `66`.<br>

### Constraints:

- 1 <= `nums.length` <= 50
- 1 <= `nums[i]` <= 1000

### <i>Concepts Used:
- Array
- Math </i>
# 3411. Maximum Subarray With Equal Products
## Question Level: Easy
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


### Examples:
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

### <i>Weekly Contest Question</i>
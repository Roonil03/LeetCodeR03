# 3932. Count K-th Roots in a Range
## Question Level: Medium
### Description:
You are given three integers l, r, and k.

An integer y is said to be a perfect kth power if there exists an integer x such that y = x<sup>k</sup>.

Return the number of integers y in the range `[l, r]` (inclusive) that are perfect k<sup>th</sup> powers.

### Examples:
#### Example 1:

Input: l = 1, r = 9, k = 3

Output: 2

Explanation:
The perfect cubes in the range `[1, 9]` are:
- 1 = 1<sup>3</sup>
- 8 = 2<sup>3</sup>  
Hence, the answer is 2.
#### Example 2:

Input: l = 8, r = 30, k = 2

Output: 3

Explanation:

The perfect squares in the range `[8, 30]` are:
- 9 = 3<sup>2</sup>
- 16 = 4<sup>2</sup>
- 25 = 5<sup>2</sup>  
Hence, the answer is 3.

### Constraints:

- 0 <= l <= r <= 10<sup>9</sup>
- 1 <= k <= 30
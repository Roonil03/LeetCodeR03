# 135. Candy
## Question Level: Hard
### Description:
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:
- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.

### Examples:
#### Example 1:

Input: ratings = `[1,0,2]`<br>
Output: 5<br>
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.<br>
#### Example 2:

Input: ratings = `[1,2,2]`<br>
Output: 4<br>
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.<br>
The third child gets 1 candy because it satisfies the above two conditions.

### Constraints:

- n == `ratings.length`
- 1 <= n <= 2 * 10^4
- 0 <= `ratings[i]` <= 2 * 10^4

### <i>Concepts Used:
- Array
- Greedy </i>
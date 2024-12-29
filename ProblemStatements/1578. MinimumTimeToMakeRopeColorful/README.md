# 1578. Minimum Time to Make Rope Colorful
## Question Level: Medium
### Description:
Alice has n balloons arranged on a rope. You are given a 0-indexed string colors where `colors[i]` is the color of the ith balloon.

Alice wants the rope to be colorful. She does not want two consecutive balloons to be of the same color, so she asks Bob for help. Bob can remove some balloons from the rope to make it colorful. You are given a 0-indexed integer array neededTime where `neededTime[i]` is the time (in seconds) that Bob needs to remove the ith balloon from the rope.

Return the minimum time Bob needs to make the rope colorful.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/12/13/ballon1.jpg"><br>
Input: colors = "`abaac`", neededTime = `[1,2,3,4,5]`<br>
Output: 3<br>
Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.<br>
Bob can remove the blue balloon at index 2. This takes 3 seconds.<br>
There are no longer two consecutive balloons of the same color. Total time = 3.<br>
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/12/13/balloon2.jpg"><br>
Input: colors = "`abc`", neededTime = `[1,2,3]`<br>
Output: 0<br>
Explanation: The rope is already colorful. Bob does not need to remove any balloons from the rope.<br>
#### Example 3:

<img src="https://assets.leetcode.com/uploads/2021/12/13/balloon3.jpg"><br>
Input: colors = "`aabaa`", neededTime = `[1,2,3,4,1]`<br>
Output: 2<br>
Explanation: Bob will remove the balloons at indices 0 and 4. Each balloons takes 1 second to remove.<br>
There are no longer two consecutive balloons of the same color. Total time = 1 + 1 = 2.<br>
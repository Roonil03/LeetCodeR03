# 2491. Divide Players Into Teams of Equal Skill
## Question Level: Medium
### Description:
You are given a positive integer array skill of even length n where skill[i] denotes the skill of the ith player. Divide the players into n / 2 teams of size 2 such that the total skill of each team is equal.<br>
The chemistry of a team is equal to the product of the skills of the players on that team.<br>
Return the sum of the chemistry of all the teams, or return -1 if there is no way to divide the players into teams such that the total skill of each team is equal.<br>


### Examples:
<b>Example 1:</b><br>
Input: skill = [3,2,5,1,3,4]<br>
Output: 22<br>
Explanation: <br>
Divide the players into the following teams: (1, 5), (2, 4), (3, 3), where each team has a total skill of 6.<br>
The sum of the chemistry of all the teams is: 1 * 5 + 2 * 4 + 3 * 3 = 5 + 8 + 9 = 22.<br>

<b>Example 2:</b><br>
Input: skill = [3,4]<br>
Output: 12<br>
Explanation: <br>
The two players form a team with a total skill of 7.<br>
The chemistry of the team is 3 * 4 = 12.<br>

<b>Example 3:</b><br>
Input: skill = [1,1,2,3]<br>
Output: -1<br>
Explanation: <br>
There is no way to divide the players into teams such that the total skill of each team is equal.<br>

### Constraints:

- 2 <= skill.length <= 10^5
- skill.length is even.
- 1 <= skill[i] <= 1000
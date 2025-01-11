# 79. Word Search
## Question Level: Medium
### Description:
Given an `m x n` grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2020/11/04/word2.jpg"><br>
Input: board = `[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, word = "`ABCCED`"<br>
Output: true
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg"><br>
Input: board = `[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, word = "`SEE`"<br>
Output: true
#### Example 3:
<img src"https://assets.leetcode.com/uploads/2020/10/15/word3.jpg"><br>
Input: board = `[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, word = "`ABCB`"<br>
Output: false

### Constraints:

- m == `board.length`
- n = `board[i].length`
- 1 <= `m`, `n` <= 6
- 1 <= `word.length` <= 15
- `board` and `word` consists of only lowercase and uppercase English letters.

### <i>Follow up: 
Could you use search pruning to make your solution faster with a larger board?

### Concepts Used:
- Array
- String
- Backtracking
- Depth-First Search
- Matrix </i>
 

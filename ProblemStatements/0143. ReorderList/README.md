# 143. Reorder List
## Question Level: Medium
### Description:
You are given the head of a singly linked-list. The list can be represented as:

L<sub>0</sub> → L<sub>1</sub> → … → L<sub>n - 1</sub> → L<sub>n</sub>
Reorder the list to be on the following form:

L<sub>0</sub> → L<sub>n</sub> → L<sub>1</sub> → L<sub>n - 1</sub> → L<sub>2</sub> → L<sub>n - 2</sub> → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg"><br>
Input: head = `[1,2,3,4]`  
Output: `[1,4,2,3]`  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg"><br>
Input: head = `[1,2,3,4,5]`  
Output: `[1,5,2,4,3]`  

### Constraints:

- The number of nodes in the list is in the range [1, 5 * 10<sup>4</sup>].
- 1 <= `Node.val` <= 1000

### <i>Concepts Used:
- Linked List
- Two Pointers
- Stack
- Recursion </i>
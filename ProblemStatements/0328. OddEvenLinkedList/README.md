# 328. Odd Even Linked List
## Question Level: Medium
### Description:
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

### Examples
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg"><br>
Input: head = `[1,2,3,4,5]`
Output: `[1,3,5,2,4]`
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg"><br>
Input: head = `[2,1,3,5,6,4,7]`
Output: `[2,3,6,7,1,5,4]`

### Constraints:

- The number of nodes in the linked list is in the range `[0, 10`<sup>`4`</sup>`]`.
- -10<sup>6</sup> <= `Node.val` <= 10<sup>6</sup>

### <i>Concepts Used:
- Linked List</i>
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if(head == None):
            return head
        size = 0
        n = head
        while(n.next):
            n = n.next
            size += 1
        size += 1
        n.next = head
        k = k % size
        if k == 0:
            n.next = None
            return head
        temp = head
        for i in range(size - k - 1):
            temp = temp.next
        res = temp.next
        temp.next = None
        return res
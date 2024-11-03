from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        def reverseLinkedList(head, k):
            prev, curr = None, head
            while k:
                next_node = curr.next
                curr.next = prev
                prev = curr
                curr = next_node
                k -= 1
            return prev
        dummy = ListNode(0)
        dummy.next = head
        group_prev = dummy

        while True:
            kth_node = group_prev
            for _ in range(k):
                kth_node = kth_node.next
                if not kth_node:
                    return dummy.next

            group_next = kth_node.next
            prev, curr = group_next, group_prev.next
            for _ in range(k):
                next_node = curr.next
                curr.next = prev
                prev = curr
                curr = next_node
            tmp = group_prev.next
            group_prev.next = kth_node
            group_prev = tmp
        return dummy.next

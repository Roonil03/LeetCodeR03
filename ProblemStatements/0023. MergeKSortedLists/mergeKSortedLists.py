from heapq import heappush, heappop
from typing import List, Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        min_heap = []              
        for i, node in enumerate(lists):
            if node:
                heappush(min_heap, (node.val, i, node))  
        dummy = ListNode()
        current = dummy        
        while min_heap:
            value, idx, node = heappop(min_heap)
            current.next = node
            current = current.next
            if node.next:
                heappush(min_heap, (node.next.val, idx, node.next))
        return dummy.next
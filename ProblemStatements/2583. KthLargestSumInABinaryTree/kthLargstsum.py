# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
from collections import deque
import heapq

class Solution:
    def kthLargestLevelSum(self, root: Optional[TreeNode], k: int) -> int:
        if not root:
            return -1

        queue = deque([(root, 0)])
        level_sums = []
        level_sums_dict = {}

        while queue:
            node, level = queue.popleft()

            if level not in level_sums_dict:
                level_sums_dict[level] = 0
            level_sums_dict[level] += node.val

            if node.left:
                queue.append((node.left, level + 1))
            if node.right:
                queue.append((node.right, level + 1))

        level_sums = list(level_sums_dict.values())

        if k > len(level_sums):
            return -1

        return heapq.nlargest(k, level_sums)[-1]
     
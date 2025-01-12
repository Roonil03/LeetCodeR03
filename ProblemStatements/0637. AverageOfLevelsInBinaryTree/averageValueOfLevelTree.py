# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        q = deque([root])
        res = []
        while q:
            sums = 0
            count = len(q)
            for _ in range(count):
                node = q.popleft()
                sums += node.val
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            avg = sums / count
            res.append(avg)
        return res
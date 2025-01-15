class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        s = []
        res = 0
        heights.append(0)
        for i, h in enumerate(heights):
            while s and heights[s[-1]] > h:
                p = heights[s.pop()]
                w = i if not s else i - s[-1] - 1
                res = max(res, p * w)
            s.append(i)
        return res

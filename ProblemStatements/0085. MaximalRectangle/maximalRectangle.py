class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        c = len(matrix[0])
        h = [0] * (c + 1)
        res = 0
        for r in matrix:
            for i in range(c):
                h[i] = h[i] + 1 if r[i] == '1' else 0
            s = [-1]
            for i in range(c + 1):
                while h[i] < h[s[-1]]:
                    x = h[s.pop()]
                    w = i - s[-1] - 1
                    res = max(res, x*w)
                s.append(i)
        return res
class Solution:
    def outerTrees(self, trees: List[List[int]]) -> List[List[int]]:
        if len(trees) == 1:
            return trees
        def orit(p, q, r):
            return (q[1] - p[1])*(r[0] - q[0]) - (q[0] - p[0]) * (r[1] - q[1])
        def bwae(p, i, q):
            return (i[0] >= p[0] and i[0] <= q[0] or i[0] <= p[0] and i[0] >= q[0]) and \
               (i[1] >= p[1] and i[1] <= q[1] or i[1] <= p[1] and i[1] >= q[1])
        trees = sorted(map(tuple, trees))
        lower = []
        for p in trees:
            while len(lower) >= 2 and orit(lower[-2], lower[-1], p) < 0:
                lower.pop()
            lower.append(p)
        upper = []
        for p in reversed(trees):
            while len(upper) >= 2 and orit(upper[-2], upper[-1], p) < 0:
                upper.pop()
            upper.append(p)
        return list(set(lower[:-1] + upper[:-1]))
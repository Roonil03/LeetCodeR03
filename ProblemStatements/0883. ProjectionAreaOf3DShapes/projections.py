class Solution:
    def projectionArea(self, grid: List[List[int]]) -> int:
         return sum(cell > 0 for row in grid for cell in row) + sum(map(max, grid + list(zip(*grid))))
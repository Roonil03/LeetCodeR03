from typing import List
import heapq

class Solution:
    def maxAverageRatio(self, classes: List[List[int]], extraStudents: int) -> float:
        def improvement(p, t):
            return (p + 1) / (t + 1) - p / t        
        heap = [(-improvement(p, t), p, t) for p, t in classes]
        heapq.heapify(heap)
        for i in range(extraStudents):
            imp, p, t = heapq.heappop(heap)
            heapq.heappush(heap, (-improvement(p + 1, t + 1), p + 1, t + 1))
        return sum(p / t for j, p, t in heap) / len(classes)
import heapq
from typing import List

class Solution:
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        min_heap = []
        currentMax = float('-inf')
        for i in range(len(nums)):
            heapq.heappush(min_heap,(nums[i][0],i,0))
            currentMax = max(currentMax,nums[i][0])
        bestRange = float('inf')
        rangeStart, rangeEnd = -1, 1
        while min_heap:
            currentMin,listIndex, elementIndex = heapq.heappop(min_heap)
            if currentMax - currentMin < bestRange:
                bestRange = currentMax - currentMin
                rangeStart = currentMin
                rangeEnd = currentMax
            if elementIndex + 1 == len(nums[listIndex]):
                break
            nextValue = nums[listIndex][elementIndex + 1]
            heapq.heappush(min_heap, (nextValue, listIndex, elementIndex+1))
            currentMax = max(currentMax,nextValue)

        return[rangeStart,rangeEnd]
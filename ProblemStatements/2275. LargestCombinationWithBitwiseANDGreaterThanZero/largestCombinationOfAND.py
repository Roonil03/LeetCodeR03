from typing import List

class Solution:
    def largestCombination(self, candidates: List[int]) -> int:
        max_bits = 32  
        bit_count = [0] * max_bits
        for num in candidates:
            for i in range(max_bits):
                if num & (1 << i):
                    bit_count[i] += 1
        return max(bit_count)
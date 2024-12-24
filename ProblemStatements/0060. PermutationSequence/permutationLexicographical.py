import math

class Solution:
    def getPermutation(self, n: int, k: int) -> str:        
        numbers = list(range(1, n + 1))
        k -= 1
        perm = []
        for i in range(n, 0, -1):
            fact = math.factorial(i - 1)
            index = k // fact
            perm.append(numbers.pop(index))
            k %= fact
        return ''.join(map(str, perm))
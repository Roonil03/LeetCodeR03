class Solution:
    def minimumSteps(self, s: str) -> int:
        wBalls = 0
        result = 0
        j = 0
        n = len(s)
        while(j<n):
            if s[j] == '0':
                result += wBalls
            wBalls += 1 if s[j] == '1' else 0
            j+= 1
        return result
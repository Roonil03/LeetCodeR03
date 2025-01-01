class Solution:
    def evaluate(self, s: str, knowledge: List[List[str]]) -> str:
        kd = {k : v for k, v in knowledge}
        res = []
        i = 0
        while i < len(s):
            if s[i] == '(':
                j = i
                while s[j] != ')':
                    j += 1
                k = s[i+1 : j]
                res.append(kd.get(k, '?'))
                i = j+1
            else:
                res.append(s[i])
                i += 1
        return ''.join(res)
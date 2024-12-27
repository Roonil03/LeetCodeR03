class Solution:
    def distinctEchoSubstrings(self, text: str) -> int:
        n = len(text)
        base = 31
        mod = 2**61 - 1
        hashed = [0] * (n + 1)
        power = [1] * (n + 1)
        for i in range(n):
            hashed[i + 1] = (hashed[i] * base + ord(text[i])) % mod
            power[i + 1] = (power[i] * base) % mod
        def getHash(l, r):
            return (hashed[r] - hashed[l] * power[r - l] % mod + mod) % mod
        res = set()
        for l in range(1, n // 2 + 1):
            for i in range(n - 2 * l + 1):
                if getHash(i, i + l) == getHash(i + l, i + 2 * l):
                    res.add(getHash(i, i + 2 * l))
        return len(res)

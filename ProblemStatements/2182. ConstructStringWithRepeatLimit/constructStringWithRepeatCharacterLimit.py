from heapq import heappop, heappush

class Solution:
    def repeatLimitedString(self, s: str, repeatLimit: int) -> str:
        d = {}
        h = []
        for ch in s:
            d[ch] = d.get(ch,0) + 1
        for ch, count in d.items():
            heappush(h, (-ord(ch), count))
        res = []
        while h:
            x, v = heappop(h)
            if res and res[-1] == chr(-x):
                if not h:
                    break
                y, w = heappop(h)
                res.append(chr(-y))
                w -= 1
                if w > 0:
                    heappush(h, (y, w))
                heappush(h, (x, v))
            else:
                k = min(repeatLimit, v)
                res.extend([chr(-x)] * k)
                v -= k
                if v > 0:
                    heappush(h, (x, v))
        return "".join(res)
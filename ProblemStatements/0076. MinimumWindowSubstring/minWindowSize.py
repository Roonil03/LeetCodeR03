class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not s or not t:
            return ""
        d = Counter(t)
        req = len(d)
        l, r = 0, 0
        formed = 0
        win = defaultdict(int)
        ans = float("inf"), None, None
        while r < len(s):
            ch = s[r]
            win[ch] += 1
            if ch in d and win[ch] == d[ch]:
                formed += 1
            while l <= r and formed == req:
                ch = s[l]
                if r - l + 1 < ans[0]:
                    ans = (r - l + 1, l, r)
                win[ch] -= 1
                if ch in d and win[ch] < d[ch]:
                    formed -= 1
                l += 1
            r += 1
        return "" if ans[0] == float("inf") else s[ans[1]: ans[2] + 1]
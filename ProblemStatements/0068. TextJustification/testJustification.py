class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        res = []
        cur = []
        num = 0
        for w in words:
            if num + len(w) + len(cur) > maxWidth:
                for i in range(maxWidth - num):
                    cur[i % (len(cur) - 1 or 1)] += ' '
                res.append(''.join(cur))
                cur, num = [], 0
            cur += [w]
            num += len(w)
        res.append(' '.join(cur).ljust(maxWidth))
        return res
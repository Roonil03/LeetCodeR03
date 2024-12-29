# Could not solve it entirely during the time of the contest, 
# it has been solved post the time limit of the contest.
# Was so close to solving it, it infuriates me that I was not able to during the contest

class Solution:
    def answerString(self, word: str, numFriends: int) -> str:
        if numFriends == 1:
            return word
        res = ""
        n = len(word)
        numFriends -= 1
        for i in range(n):
            u = numFriends - i
            if n - max(0,u) <= i:
                continue
            w = word[i:n - max(0,u)]
            if not res:
                res = w
            else:
                res = max(res, w)
        return res
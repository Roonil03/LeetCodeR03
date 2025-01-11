class Solution:
    # def maxTaskAssign(self, tasks: List[int], workers: List[int], pills: int, strength: int) -> int:
    #     tasks.sort()
    #     workers.sort()        
    #     def comp(k):
    #         q = deque(workers[len(workers) - k:])
    #         left = pills            
    #         for i in range(k - 1, -1, -1):
    #             if not q or q[-1] + strength < tasks[i]:
    #                 return False                    
    #             if q[-1] >= tasks[i]:
    #                 q.pop()
    #             elif left > 0:
    #                 s = deque()
    #                 f = False                    
    #                 while q:
    #                     w = q.popleft()
    #                     if w + strength >= tasks[i]:
    #                         left -= 1
    #                         f = True
    #                         while s:
    #                             q.appendleft(s.pop())
    #                         break
    #                     s.append(w)                    
    #                 if not f:
    #                     return False                        
    #                 while s:
    #                     q.appendleft(s.pop())
    #             else:
    #                 return False                    
    #         return True        
    #     l, r = 0, min(len(tasks), len(workers))
    #     while l < r:
    #         mid = (l + r + 1) // 2
    #         if comp(mid):
    #             l = mid
    #         else:
    #             r = mid - 1                
    #     return l
    def maxTaskAssign(self, tasks: List[int], workers: List[int], pills: int, strength: int) -> int:
        tasks.sort()
        workers.sort()
        
        def fn(k, p=pills): 
            ww = workers[-k:]
            for t in reversed(tasks[:k]): 
                if t <= ww[-1]: ww.pop()
                elif t <= ww[-1] + strength and p: 
                    p -= 1
                    i = bisect_left(ww, t - strength)
                    ww.pop(i)
                else: return False 
            return True 
          
        lo, hi = 0, min(len(tasks), len(workers))
        while lo < hi: 
            mid = lo + hi + 1 >> 1
            if fn(mid): lo = mid
            else: hi = mid - 1
        return lo 
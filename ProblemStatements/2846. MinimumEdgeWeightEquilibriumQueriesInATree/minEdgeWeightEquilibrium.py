class Solution:
    def minOperationsQueries(self, n: int, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        adj = [[] for _ in range(n)]
        for u, v, w in edges: 
            adj[u].append((v, w))
            adj[v].append((u, w))        
        logN = int(log2(n)) + 2
        anc = [[-1]*logN for _ in range(n)]
        freq = [[0]*27 for _ in range(n)]
        dep = [-1]*n 
        stk = [(0, -1, 0)]
        while stk: 
            nd, par, d = stk.pop()
            dep[nd] = d 
            for nei, wt in adj[nd]: 
                if nei != par: 
                    anc[nei][0] = nd
                    freq[nei][:] = freq[nd][:]
                    freq[nei][wt] += 1
                    for j in range(1, logN): 
                        if anc[nei][j-1] == -1: 
                            break 
                        anc[nei][j] = anc[anc[nei][j-1]][j-1]
                    stk.append((nei, nd, d+1))   
        def lca(u, v): 
            if dep[u] > dep[v]: 
                u, v = v, u
            for i in range(logN): 
                if dep[v] - dep[u] & (1 << i): 
                    v = anc[v][i]
            if u == v: 
                return u 
            for i in range(logN-1, -1, -1): 
                if anc[u][i] != anc[v][i]: 
                    u, v = anc[u][i], anc[v][i]
            return anc[u][0]
        res = []
        for u, v in queries: 
            k = lca(u, v)
            pathFreq = [freq[u][w] + freq[v][w] - 2*freq[k][w] for w in range(27)]
            res.append(sum(pathFreq) - max(pathFreq))
        
        return res
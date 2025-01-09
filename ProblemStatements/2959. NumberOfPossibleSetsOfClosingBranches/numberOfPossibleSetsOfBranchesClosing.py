class Solution:
    def numberOfSets(self, n: int, maxDistance: int, roads: List[List[int]]) -> int:
        inf = float('inf')
        graph = [[inf] * n for _ in range(n)]        
        for u, v, w in roads:
            if graph[u][v] > w:
                graph[u][v] = w
                graph[v][u] = w        
        for i in range(n):
            graph[i][i] = 0
        def floyd_warshall(mask):
            dist = [[inf] * n for _ in range(n)]
            for i in range(n):
                if mask & (1 << i):
                    for j in range(n):
                        if mask & (1 << j):
                            dist[i][j] = graph[i][j]
            for k in range(n):
                if mask & (1 << k):
                    for i in range(n):
                        if mask & (1 << i):
                            for j in range(n):
                                if mask & (1 << j):
                                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
            return dist
        def is_valid_set(mask):
            dist = floyd_warshall(mask)
            for i in range(n):
                if mask & (1 << i):
                    for j in range(n):
                        if mask & (1 << j) and dist[i][j] > maxDistance:
                            return False
            return True
        res = 1
        for mask in range(1, 1 << n):
            if is_valid_set(mask):
                res += 1
        return res
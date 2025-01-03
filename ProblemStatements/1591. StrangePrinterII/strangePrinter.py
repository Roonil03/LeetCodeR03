class Solution:
    def isPrintable(self, targetGrid):
        graph = defaultdict(list)
        inDegree = [0] * 61        
        for color in range(1, 61):
            self.search(targetGrid, color, graph, inDegree)
        zeros = deque()
        seen = set()        
        for i in range(len(inDegree)):
            if inDegree[i] == 0:
                zeros.append(i)        
        while zeros:
            cur = zeros.popleft()
            if cur in seen:
                continue
            seen.add(cur)            
            for neighbor in graph[cur]:
                inDegree[neighbor] -= 1
                if inDegree[neighbor] == 0:
                    zeros.append(neighbor)        
        return len(seen) == 61

    def search(self, grid, color, graph, inDegree):
        minX, minY = float('inf'), float('inf')
        maxX, maxY = float('-inf'), float('-inf')        
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == color:
                    minX = min(minX, i)
                    maxX = max(maxX, i)
                    minY = min(minY, j)
                    maxY = max(maxY, j)        
        if minX == float('inf'):
            return        
        for i in range(minX, maxX + 1):
            for j in range(minY, maxY + 1):
                if grid[i][j] != color:
                    graph[grid[i][j]].append(color)
                    inDegree[color] += 1
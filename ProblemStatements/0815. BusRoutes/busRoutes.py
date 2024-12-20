class Solution:
    def numBusesToDestination(self, routes: List[List[int]], source: int, target: int) -> int:
        if source == target:
            return 0
        stop = defaultdict(set)
        for i, route in enumerate(routes):
            for s in route:
                stop[s].add(i)
        visited = set()
        visRoutes = set()
        q = deque([source])
        bus = 0
        while q:
            bus += 1
            for _ in range(len(q)):
                currStop = q.popleft()
                for r in stop[currStop]:
                    if r in visRoutes:
                        continue
                    visRoutes.add(r)
                    for s in routes[r]:
                        if s == target:
                            return bus
                        if s not in visited:
                            visited.add(s)
                            q.append(s)
        return -1
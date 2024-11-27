public class Solution {
    public int[] ShortestDistanceAfterQueries(int n, int[][] queries) {
        var adjList = new List<HashSet<int>>(n);
        for (int i = 0; i < n; i++) {
            adjList.Add(new HashSet<int>());
        }
        for (int i = 0; i < n - 1; i++) {
            adjList[i].Add(i + 1);
        }
        var result = new int[queries.Length];
        for (int q = 0; q < queries.Length; q++) {
            int u = queries[q][0];
            int v = queries[q][1];
            if (!adjList[u].Contains(v)) {
                adjList[u].Add(v);
            }
            result[q] = BFS(n, adjList);
        }
        return result;
    }

    private int BFS(int n, List<HashSet<int>> adjList) {
        var visited = new bool[n];
        var queue = new Queue<int>();
        queue.Enqueue(0);
        visited[0] = true;
        int layers = 0;
        while (queue.Count > 0) {
            int size = queue.Count;
            for (int i = 0; i < size; i++) {
                int current = queue.Dequeue();
                if (current == n - 1) {
                    return layers;
                }
                foreach (var neighbor in adjList[current]) {
                    if (!visited[neighbor]) {
                        visited[neighbor] = true;
                        queue.Enqueue(neighbor);
                    }
                }
            }
            layers++;
        }
        return -1;
    }
}

// Could not solve during contest, has been solved post the contest time.

public class Solution {
    public int MinMaxWeight(int n, int[][] edges, int threshold) {
        int l = 0, r = 1000000;
        while (l <= r) {
            int mid = (l + r) / 2;
            List<List<int>> e = new List<List<int>>(n);
            for (int i = 0; i < n; i++) {
                e.Add(new List<int>());
            }
            foreach (var edge in edges) {
                int u = edge[0], v = edge[1], w = edge[2];
                if (w > mid) continue;
                e[v].Add(u);
            }
            bool[] vis = new bool[n];
            Queue<int> q = new Queue<int>();
            q.Enqueue(0);
            vis[0] = true;

            while (q.Count > 0) {
                int x = q.Dequeue();
                foreach (var y in e[x]) {
                    if (vis[y]) continue;
                    vis[y] = true;
                    q.Enqueue(y);
                }
            }
            bool flag = true;
            for (int i = 0; i < n; i++) {
                if (!vis[i]) {
                    flag = false;
                    break;
                }
            }
            if (!flag)
                l = mid + 1;
            else
                r = mid - 1;
        }
        if (l > 1000000) return -1;
        return l;
    }
}

// This program utilizes a combination of binary search and breadth-first search (BFS) to solve the problem 
// of finding the minimum maximum edge weight in a graph that allows all nodes to be reachable within a 
// given threshold. The binary search is used to efficiently narrow down the range of possible maximum edge 
// weights, starting from 0 to 1,000,000. For each midpoint value (mid), the graph is reconstructed by 
// including only those edges whose weights are less than or equal to mid. The BFS is then used to check the 
// connectivity of the graph, starting from node 0 and ensuring all nodes are reachable. If all nodes are 
// reachable, the search continues with a lower mid value to find a potentially smaller maximum edge weight. 
// If not all nodes are reachable, the search continues with a higher mid value. This approach ensures that 
// the solution is both efficient and effective in determining the minimum maximum edge weight that satisfies 
// the connectivity requirement.
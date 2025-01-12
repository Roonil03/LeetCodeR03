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
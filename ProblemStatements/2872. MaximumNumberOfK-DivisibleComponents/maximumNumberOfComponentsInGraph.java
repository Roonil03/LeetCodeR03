class Solution {
    public int maxKDivisibleComponents(int n, int[][] edges, int[] values, int k) {
        int[] sums = new int[n];
        int[] ans = {0};
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i < n; i++) graph.add(new ArrayList<>());
        for (int[] e : edges) {
            graph.get(e[0]).add(e[1]);
            graph.get(e[1]).add(e[0]);
        }
        dfs(0, -1, graph, values, sums, k, ans);
        return ans[0];
    }
    
    int dfs(int node, int parent, List<List<Integer>> graph, int[] values, int[] sums, int k, int[] ans) {
        sums[node] = values[node];
        for (int n : graph.get(node)) {
            if (n == parent)
            {
                continue;
            }
            sums[node] += dfs(n, node, graph, values, sums, k, ans);
        }
        if (sums[node] % k == 0)
        {
            ans[0]++;
            return 0;
        }
        return sums[node] % k;
    }
}

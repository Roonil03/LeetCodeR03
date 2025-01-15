public class Solution
{
    private List<int>[] graph;
    private bool[] seen;
    private int bob;
    private int[] amount;

    public int MostProfitablePath(int[][] edges, int bob, int[] amount)
    {
        int n = edges.Length + 1;
        this.bob = bob;
        this.amount = amount;
        graph = new List<int>[n];
        for (int i = 0; i < n; i++)
        {
            graph[i] = new List<int>();
        }
        foreach (var edge in edges)
        {
            graph[edge[0]].Add(edge[1]);
            graph[edge[1]].Add(edge[0]);
        }
        seen = new bool[n];
        return Dfs(0, 0).Item1;
    }

    private Tuple<int, int> Dfs(int node, int depth)
    {
        seen[node] = true;
        int res = int.MinValue;
        int bobDepth = (node == bob) ? 0 : seen.Length;        
        foreach (int neighbor in graph[node])
        {
            if (seen[neighbor]) continue;
            var (curRes, curBobDepth) = Dfs(neighbor, depth + 1);
            res = Math.Max(res, curRes);
            bobDepth = Math.Min(bobDepth, curBobDepth);
        }        
        if (res == int.MinValue) res = 0;
        if (depth == bobDepth) res += amount[node] / 2;
        if (depth < bobDepth) res += amount[node];
        return new Tuple<int, int>(res, bobDepth + 1);
    }
}
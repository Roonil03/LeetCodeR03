/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution {
    public int MinimumOperations(TreeNode root) {
        var q = new Queue<TreeNode>();
        q.Enqueue(root);
        int ops = 0;
        while (q.Count > 0) {
            int lvl = q.Count;
            var c = new List<int>();
            for (int i = 0; i < lvl; i++) {
                var node = q.Dequeue();
                c.Add(node.val);
                if (node.left != null) {
                    q.Enqueue(node.left);
                }
                if (node.right != null) {
                    q.Enqueue(node.right);
                }
            }
            ops += swaps(c);
        }
        return ops;
    }

    int swaps(List<int> c) {
        int n = c.Count;
        var sor = c.Select((c, id) => new { c, id }).OrderBy(x => x.c).ToList();
        var visited = new bool[n];
        int swaps = 0;
        for (int i = 0; i < n; i++) {
            if (visited[i] || sor[i].id == i) {
                continue;
            }
            int size = 0;
            int j = i;
            while (!visited[j]) {
                visited[j] = true;
                j = sor[j].id;
                size++;
            }
            if (size > 1) {
                swaps += size - 1;
            }
        }
        return swaps;
    }
}

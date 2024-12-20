/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode reverseOddLevels(TreeNode root) {
        if (root == null)
        {
            return null;
        }
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        boolean odd = false;
        while (!q.isEmpty()) {
            int size = q.size();
            List<TreeNode> levelNodes = new ArrayList<>();
            for (int i = 0; i < size; i++) {
                TreeNode node = q.poll();
                if (odd)
                {
                    levelNodes.add(node);
                }
                if (node.left != null) {
                    q.offer(node.left);
                    q.offer(node.right);
                }
            }
            if (odd) 
            {
                int l = 0, r = levelNodes.size() - 1;
                while (l < r) {
                    int temp = levelNodes.get(l).val;
                    levelNodes.get(l).val = levelNodes.get(r).val;
                    levelNodes.get(r).val = temp;
                    l++;
                    r--;
                }
            }
            odd = !odd;
        }
        return root;
    }
}
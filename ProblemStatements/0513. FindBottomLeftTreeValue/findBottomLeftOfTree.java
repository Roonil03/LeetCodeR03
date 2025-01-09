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
    public int findBottomLeftValue(TreeNode root) {
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        int res = root.val;
        while(!q.isEmpty())
        {
            int ls = q.size();
            for(int i = 0; i < ls; i++)
            {
                TreeNode temp = q.poll();
                if(i==0)
                {
                    res = temp.val;
                }
                if(temp.left != null)
                {
                    q.offer(temp.left);
                }
                if(temp.right != null)
                {
                    q.offer(temp.right);
                }
            }
        }
        return res;
    }
}
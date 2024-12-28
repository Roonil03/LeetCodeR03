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
    public List<Integer> largestValues(TreeNode root) {
       List<Integer> res = new ArrayList<>();
       if (root == null) {
           return res;
       }
       Queue<TreeNode> q = new LinkedList<>();
       q.offer(root);
       while (!q.isEmpty()) {
           int size = q.size();
           int max = Integer.MIN_VALUE;
           for (int i = 0; i < size; i++) {
               TreeNode temp = q.poll();
               if (temp.val > max) {
                   max = temp.val;
               }
               if (temp.left != null) {
                   q.offer(temp.left);
               }
               if (temp.right != null) {
                   q.offer(temp.right);
               }
           }
           res.add(max);
       }
       return res;
   }
}
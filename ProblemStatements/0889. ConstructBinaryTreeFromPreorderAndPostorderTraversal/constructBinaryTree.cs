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
    public TreeNode ConstructFromPrePost(int[] preorder, int[] postorder) {
        int pre = 0, post = 0;        
        TreeNode BuildTree() {
            TreeNode root = new TreeNode(preorder[pre++]);
            if (root.val != postorder[post]) {
                root.left = BuildTree();
            }
            if (root.val != postorder[post]) {
                root.right = BuildTree();
            }
            post++;
            return root;
        }        
        return BuildTree();
    }
}

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
class BSTIterator {

    Stack<TreeNode> s;

    public BSTIterator(TreeNode root) {
        s = new Stack<>();
        TreeNode c = root;
        while(c != null){
            s.push(c);
            if (c.left != null)
                c = c.left;
            else
                break;
        }
    }
    
    public boolean hasNext() {
        return !s.isEmpty();
    }
    
    public int next() {
        TreeNode n = s.pop();
        TreeNode c = n;
        if(c.right != null){
            c = c.right;
            while(c != null){
                s.push(c);
                if (c.left != null){
                    c = c.left;
                } else{
                    break;
                }
            }
        }
        return n.val;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * BSTIterator obj = new BSTIterator(root);
 * int param_1 = obj.next();
 * boolean param_2 = obj.hasNext();
 */
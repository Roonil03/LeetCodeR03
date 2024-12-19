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

import java.util.*;

class Solution {
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> ls = new LinkedList<>();
        if(root == null)
        {
            return ls;
        }
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        while(!q.isEmpty())
        {
            int size = q.size();
            List<Integer> temp = new ArrayList<>();
            for(int i = 0; i< size; i++)
            {
                TreeNode t = q.poll();
                temp.add(t.val);
                if(t.left != null)
                {
                    q.offer(t.left);
                }
                if(t.right != null)
                {
                    q.offer(t.right);
                }
            }
            ls.add(0,temp);
        }
        return ls;
    }
}
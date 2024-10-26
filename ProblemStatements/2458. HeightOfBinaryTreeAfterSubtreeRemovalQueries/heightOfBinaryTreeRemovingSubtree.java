
//Definition for a binary tree node.
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
     TreeNode(int val) { this.val = val; }
     TreeNode(int val, TreeNode left, TreeNode right) {
         this.val = val;
         this.left = left;
         this.right = right;
    }
}

class Solution{
   int[] nodeDepth;
   int[] nodeLevel;
   int[] primaryMax;
   int[] secondaryMax;

   public int height(TreeNode root, int level) {
       if (root == null) {
           return 0; 
       }
       
       nodeLevel[root.val] = level; 
       nodeDepth[root.val] = 1 + Math.max(height(root.left, level + 1), height(root.right, level + 1));
       
       if (primaryMax[level] < nodeDepth[root.val]) {
           secondaryMax[level] = primaryMax[level];
           primaryMax[level] = nodeDepth[root.val];
       } else if (secondaryMax[level] < nodeDepth[root.val]) {
           secondaryMax[level] = nodeDepth[root.val]; 
       }
       
       return nodeDepth[root.val];
   }   
   
   public int[] treeQueries(TreeNode root, int[] queries) {
       nodeDepth = new int[100001];
       nodeLevel = new int[100001];
       primaryMax = new int[100001];
       secondaryMax = new int[100001];
       
       height(root, 0);
       
       for (int i = 0; i < queries.length; i++) {
           int q = queries[i];
           int level = nodeLevel[q];
           queries[i] = (primaryMax[level] == nodeDepth[q] ? secondaryMax[level] : primaryMax[level]) + level - 1;  
       }
       
       return queries; 
   }
}
 
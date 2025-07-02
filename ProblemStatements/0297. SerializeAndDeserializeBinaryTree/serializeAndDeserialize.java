/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        helper1(root, sb);
        return sb.toString();
    }

    private void helper1(TreeNode node, StringBuilder sb){
        if(node == null){
            sb.append("#,");
            return;
        }
        sb.append(node.val).append(",");
        helper1(node.left, sb);
        helper1(node.right, sb);
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        String[] tok = data.split(",");
        Queue<String> q = new LinkedList<>(Arrays.asList(tok));
        return helper2(q);
    }

    private TreeNode helper2(Queue<String> q){
        if(q.isEmpty()){
            return null;
        }
        String tok = q.poll();
        if(tok.equals("#")){
            return null;
        }
        TreeNode node = new TreeNode(Integer.parseInt(tok));
        node.left = helper2(q);
        node.right = helper2(q);
        return node;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec ser = new Codec();
// Codec deser = new Codec();
// TreeNode ans = deser.deserialize(ser.serialize(root));
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
        if(root == null)
            return ".,";
        return root.val + "," + serialize(root.left)+serialize(root.right); 
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        return helper(new StringBuilder(data));
    }
    public TreeNode helper(StringBuilder data) {
        String value = data.substring(0,data.indexOf(","));
        if(value.equals(".")){
            data = data.delete(0, data.indexOf(",")+1);
            return null;
        }
        TreeNode n = new TreeNode(Integer.valueOf(value));
        data = data.delete(0, data.indexOf(",")+1);
        n.left = helper(data);
        n.right = helper(data);
        return n;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
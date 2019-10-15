/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        if(root == null)
            return new ArrayList<>();
        List<List<Integer>> list = new ArrayList<>();
        dfs(root, list, -1);  
        return list;
    }
    
    public void dfs(TreeNode node, List<List<Integer>> list, int prevDepth) {
        prevDepth++;
        if(list.size() < prevDepth+1)
            list.add(new ArrayList<>());
        list.get(prevDepth).add(node.val);
        if(node.left != null)
            dfs(node.left, list, prevDepth);
        if(node.right != null)
            dfs(node.right, list, prevDepth);
    }
}
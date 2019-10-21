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
    int deepestLevel = 0;
    int[] count =  new int[1000];
    public int maxLevelSum(TreeNode root) {
        dfs(root, 0);
        int max=0;
        int maxIndex = 0;
        for(int i = 0; i <= deepestLevel; i++){
            if(max < count[i]){
                max = count[i];
                maxIndex = i;
            }
        }
        return maxIndex+1;
    }
    
    public void dfs(TreeNode node, int level){
        if(node==null)
            return;
        if(level > deepestLevel)
            deepestLevel = level;
        count[level]+=node.val;
        dfs(node.left, level+1);
        dfs(node.right, level+1);
    }
}
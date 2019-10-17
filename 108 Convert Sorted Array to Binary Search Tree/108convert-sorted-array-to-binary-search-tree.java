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
    public TreeNode sortedArrayToBST(int[] nums) {
        
        if(nums == null || nums.length == 0)
            return null;
        
        return sortedArrayToBST(nums, 0, nums.length - 1);
    }
    
    public TreeNode sortedArrayToBST(int[] nums, int start, int end){
        TreeNode root = null;
        if(start > end)
            return null;
        else{
            int mid = (start + end) / 2;
            root = new TreeNode(nums[mid]);
            root.left = sortedArrayToBST(nums, start, mid - 1);
            root.right = sortedArrayToBST(nums, mid+1, end);
        }
        
        return root;
    }
}
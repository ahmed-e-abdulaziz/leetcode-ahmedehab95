/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := []int{}

    if root.Left != nil {
        res = append(res, inorderTraversal(root.Left)...)
    }
    res = append(res, root.Val)
    if root.Right != nil {
        res = append(res, inorderTraversal(root.Right)...)
    }
    return res
}
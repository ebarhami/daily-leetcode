/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    return traverse(root, &sum)
}

func traverse(root *TreeNode, sum *int) *TreeNode {
    if root != nil {
        traverse(root.Right, sum)
        *sum += root.Val
        root.Val = *sum
        traverse(root.Left, sum)
    }
    return root
}

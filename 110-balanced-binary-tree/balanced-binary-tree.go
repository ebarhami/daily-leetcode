/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    leftDepth := depth(root.Left)
    rightDepth := depth(root.Right)

    return diff(leftDepth, rightDepth) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func diff(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
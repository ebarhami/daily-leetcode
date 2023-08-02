/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    return traverse(root, root.Val, root.Val)
}

func traverse(root *TreeNode, max, min int) int {
    if root == nil {return max-min}
    if root.Val > max {max = root.Val}
    if root.Val < min {min = root.Val}

    l := traverse(root.Left, max, min)
    r := traverse(root.Right, max, min)

    return maks(l, r)
}

func maks(a, b int) int {
    if a > b {return a}
    return b
}


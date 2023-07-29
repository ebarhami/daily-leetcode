/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {return true}
    if root == nil || subRoot == nil {return false}

    if root.Val == subRoot.Val {
        if isSame(root, subRoot) {
            return true
        }
    }
    left := isSubtree(root.Left, subRoot)
    right := isSubtree(root.Right, subRoot)

    return left || right
}

func isSame(a, b *TreeNode) bool {
    if a == nil && b == nil {return true}
    if a == nil || b == nil {return false}

    check := a.Val == b.Val

    return check && isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}
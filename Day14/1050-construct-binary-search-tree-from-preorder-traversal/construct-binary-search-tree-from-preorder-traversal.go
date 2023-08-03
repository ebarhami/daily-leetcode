/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {return nil}

    node := &TreeNode{Val:preorder[0]}

    idx := 0
    for idx < len(preorder) && preorder[idx] <= preorder[0] {
        idx++
    }
    node.Left = bstFromPreorder(preorder[1:idx])
    node.Right = bstFromPreorder(preorder[idx:])

    return node
}
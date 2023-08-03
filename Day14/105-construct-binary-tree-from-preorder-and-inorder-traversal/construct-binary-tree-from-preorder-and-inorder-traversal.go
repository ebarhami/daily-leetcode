/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    idx, val := 0, preorder[0]
    for inorder[idx] != val {
        idx++
    }

    node := &TreeNode{
        Val: preorder[0],
    }
    node.Left = buildTree(preorder[1:idx+1], inorder[:idx])
    node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
    return node
}

/*
preorder = root + left + right
inorder = left + root + right

*/
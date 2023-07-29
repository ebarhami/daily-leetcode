/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }
    root := &TreeNode{}
    var left1, left2, right1, right2 *TreeNode
    
    if root1 != nil {
        root.Val += root1.Val
        left1, right1 = root1.Left, root1.Right
    }
    if root2 != nil {
        root.Val += root2.Val
        left2, right2 = root2.Left, root2.Right
    }

    root.Left = mergeTrees(left1, left2)
    root.Right = mergeTrees(right1, right2)

    return root
}
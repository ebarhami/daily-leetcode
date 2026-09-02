/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)

    var traverse func(root *TreeNode) 

    traverse = func(root *TreeNode) {
        if root == nil {
            return
        }

        traverse(root.Left)
        traverse(root.Right)
        result = append(result, root.Val)
    }

    traverse(root)
    return result
}
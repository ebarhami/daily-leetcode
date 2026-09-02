/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    answer := make([]int, 0)
    var traverse func(root *TreeNode)
    traverse = func(root *TreeNode) {
        if root == nil {
            return
        }
        traverse(root.Left)
        answer = append(answer, root.Val)
        traverse(root.Right)
    }
    traverse(root)

    return answer
}
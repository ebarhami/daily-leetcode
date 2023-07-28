/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}

    answer := make([]int, 0)
    answer = append(answer, postorderTraversal(root.Left)...)
    answer = append(answer, postorderTraversal(root.Right)...)
    answer = append(answer, root.Val)

    return answer
}
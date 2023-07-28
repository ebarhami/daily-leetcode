/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}

    answer := make([]int, 0)
    answer = append(answer, root.Val)
    answer = append(answer, preorderTraversal(root.Left)...)
    answer = append(answer, preorderTraversal(root.Right)...)

    return answer
}


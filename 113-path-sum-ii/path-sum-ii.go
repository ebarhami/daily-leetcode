/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    var traverse func(root *TreeNode, remaining int) 
    traverse = func(root *TreeNode, remaining int) {
        if root == nil {
            return
        }

        path = append(path, root.Val)
        remaining -= root.Val
        if root.Left == nil && root.Right == nil && remaining == 0 {
            result = append(result, append([]int{}, path...))
        }

        traverse(root.Left, remaining)
        traverse(root.Right, remaining)

        path = path[:len(path)-1]
    }

    traverse(root, targetSum)

    return result
}
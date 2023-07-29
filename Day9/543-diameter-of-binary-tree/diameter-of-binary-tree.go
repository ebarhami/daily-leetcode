/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxDepth := make(map[*TreeNode]int)
    
    return findAnswer(root, maxDepth)
}

func findAnswer(root *TreeNode, maxDepth map[*TreeNode]int) int {
    if root == nil {return 0}

    if _, ok := maxDepth[root.Left]; !ok {
        maxDepth[root.Left] = maxDepthF(root.Left)
    }
    if _, ok := maxDepth[root.Right]; !ok {
        maxDepth[root.Right] = maxDepthF(root.Right)
    }

    maxNow := maxDepth[root.Left] + maxDepth[root.Right]
    maxDepth[root] = maxNow

    return max(max(maxNow, findAnswer(root.Left, maxDepth)), findAnswer(root.Right, maxDepth))
}

func maxDepthF(root *TreeNode) int {
    if root == nil {return 0}
    return 1 + max(maxDepthF(root.Left), maxDepthF(root.Right))
}

func max(a, b int) int {
    if a > b {return a}
    return b
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    answer := make([]int, 0)
    rightestByLevel := make(map[int]int)
    traverse(&answer, rightestByLevel, root, 0)

    return answer
}

func traverse(answer *[]int, rightestByLevel map[int]int, root *TreeNode, level int) {
    if root == nil {return}
    if _, ok := rightestByLevel[level]; !ok {
        (*answer) = append(*answer, root.Val)
        rightestByLevel[level] = root.Val
    }

    traverse(answer, rightestByLevel, root.Right, level+1)
    traverse(answer, rightestByLevel, root.Left, level+1)

    return
}
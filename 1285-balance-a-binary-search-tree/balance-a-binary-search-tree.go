/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    sortedArray := make([]int, 0)

    var traverse func(root *TreeNode)
    traverse = func(root *TreeNode) {
        if root == nil {
            return
        }
        traverse(root.Left)
        sortedArray = append(sortedArray, root.Val)
        traverse(root.Right)
    }
    traverse(root)

    return builtBST(0, len(sortedArray)-1, sortedArray)
}

func builtBST(l, r int, array []int) *TreeNode {
    if l > r {
        return nil
    }
    mid := l + (r-l)/2
    leftTree := builtBST(l, mid-1, array)
    rightTree := builtBST(mid+1, r, array)

    return &TreeNode{
        Val: array[mid],
        Left: leftTree,
        Right: rightTree,
    }
}
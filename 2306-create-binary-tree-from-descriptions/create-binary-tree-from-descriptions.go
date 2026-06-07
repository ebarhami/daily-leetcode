/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodeMap := make(map[int]*TreeNode)
    parents := make(map[int]bool)

    for _, desc := range descriptions {
        par, child, isLeft := desc[0], desc[1], desc[2] == 1

        if _, exist := parents[par]; !exist {
            parents[par] = true
        }


            parents[child] = false
        
        
        parNode, exist := nodeMap[par]
        childNode, childExist := nodeMap[child]
        if !exist {
            parNode = &TreeNode{
                Val: par,
            }
        }
        if !childExist {
            childNode = &TreeNode{
                Val: child,
            }
        }
        
        if isLeft {
            parNode.Left = childNode
        } else {
            parNode.Right = childNode
        }
        nodeMap[par] = parNode
        nodeMap[child] = childNode
    }

    for key, val := range parents {
        if val {
            return nodeMap[key]
        }
    }

    return nil
}
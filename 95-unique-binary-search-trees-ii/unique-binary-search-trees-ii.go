/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {return nil}

    return helper(1,n)
}

func helper(min, max int) []*TreeNode {
    if min > max { return nil}
    if min == max { return []*TreeNode{{Val:min}} }

    answer := make([]*TreeNode, 0)
    for i:=min;i<=max;i++{
        left := helper(min, i-1)
        right := helper(i+1, max)
        if left == nil {
            for _, r := range right {
                answer = append(answer, &TreeNode{Val: i, Right: r})
            }
        } else if right == nil {
            for _, l := range left {
                answer = append(answer, &TreeNode{Val: i, Left: l})
            }
        } else {
            for _, l := range left {
                for _, r := range right {
                    answer = append(answer, &TreeNode{Val: i, Left: l, Right: r})
                }
            }
        }   
    }
    return answer
}
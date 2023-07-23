/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    cache := make(map[int][]*TreeNode)
    return generate(n, cache)
}

func generate(n int, cache map[int][]*TreeNode) []*TreeNode {
    if n % 2 == 0 {return nil}

    if cache[n] != nil {
        return cache[n]
    }

    if n == 1 {
        base := &TreeNode{Val:0}
        cache[n] = []*TreeNode{base}
        return cache[n]
    }

    answer := make([]*TreeNode, 0)
    for i:=1;i<n;i+=2{
        for _, l := range generate(i, cache) {
            for _, r := range generate(n-1-i, cache) {
                answer = append(answer, &TreeNode{0, l, r})
            }
        }
    }

    cache[n] = answer
    return answer
}
